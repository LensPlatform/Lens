package service

import (
	"context"
	"database/sql"
	"errors"
	"unsafe"

	"github.com/go-kit/kit/metrics"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"

	"github.com/LensPlatform/Lens/src/pkg/database"
	"github.com/LensPlatform/Lens/src/pkg/helper"
	model "github.com/LensPlatform/Lens/src/pkg/models"
)

// Service is a CRUD interface definition for the user microservice
type Service interface {
	// CreateUser effectively creates/adds a user object to the backend data store
	// if it doesm't already exist.
	CreateUser(ctx context.Context, user model.User) (err error)
	// GetUserById queries the backend datastore for user objects based on a
	// passed in user id parameter.
	GetUserById(ctx context.Context, id string) (user model.User, err error)
	// GetUserByEmail queries the backend datastore for user objects based on a
	// passed in user email parameter.
	GetUserByEmail(ctx context.Context, email string) (user model.User, err error)
	// GetUserByUsername queries the backend datastore for user objects based on a
	// passed in user username parameter.
	GetUserByUsername(ctx context.Context, username string) (user model.User, err error)
	// LogIn Checks if a user object exists in the backend datastore, performs some password checks,
	// and attempts to log a given user into the system
	LogIn(ctx context.Context, username, password string) (user model.User, err error)
}

// Counters is a type encompassing metrics for API definitions
// associated with the user microservice
type Counters struct {
	CreateUserRequest           metrics.Counter
	SuccessfulCreateUserRequest metrics.Counter
	FailedCreateUserRequest     metrics.Counter
	GetUserRequest              metrics.Counter
	SuccessfulGetUserRequest    metrics.Counter
	FailedGetUserRequest        metrics.Counter
	SuccessfulLogInRequest      metrics.Counter
	FailedLogInRequest          metrics.Counter
	Duration                    metrics.Histogram
}

var validate = validator.New()

// New returns a basic Service with all of the expected middlewares wired in.
func New(logger *zap.Logger, db *gorm.DB, amqpProducer Queue, amqpConsumer Queue, counters Counters) Service {
	var svc Service
	{
		svc = NewBasicService(db, logger, amqpProducer, amqpConsumer)
		svc = LoggingMiddleware(logger)(svc)
		svc = InstrumentingMiddleware(counters)(svc)
	}
	return svc
}

// NewBasicService returns a naïve, stateless implementation of Service.
func NewBasicService(db *gorm.DB, logger *zap.Logger, amqpProducer Queue, amqpConsumer Queue) Service {
	return basicService{logger: logger, database: database.NewDatabase(db),
		ConsumerQueues: amqpConsumer, ProducerQueues: amqpProducer}
}

// basicService is a type witholding references to logging, the backend datastroe,
// as well as queues (producer and consumer)
type basicService struct {
	logger         *zap.Logger
	database       *database.Database
	ConsumerQueues Queue
	ProducerQueues Queue
}

func (s basicService) LogIn(ctx context.Context, username, password string) (user model.User, err error) {
	if username == "" {
		s.logger.Error(helper.ErrNoUsernameProvided.Error())
		return user, helper.ErrNoUsernameProvided
	}

	if password == "" {
		s.logger.Error(helper.ErrNoPasswordProvided.Error())
		return user, helper.ErrNoPasswordProvided
	}

	// check if user exists in the database
	user, err = s.database.GetUserByUsername(username)

	if err != nil {
		if err == sql.ErrNoRows {
			s.logger.Error(helper.ErrInvalidUsernameProvided.Error())
			return user, helper.ErrInvalidUsernameProvided
		}

		return user, err
	}

	s.logger.Info("Password", zap.String("password", user.PassWord))

	// check if passwords match
	isEqual := s.comparePasswords(user.PassWord, []byte(password))

	if !isEqual {
		s.logger.Error(helper.ErrInvalidPasswordProvided.Error())
		return model.User{}, helper.ErrInvalidPasswordProvided
	}

	return user, nil
}

func (s basicService) GetUserById(ctx context.Context, id string) (user model.User, err error) {
	return s.getUserFromQueryParam(ctx, database.GetUserByIdQuery, id)
}

func (s basicService) GetUserByEmail(ctx context.Context, email string) (user model.User, err error) {
	return s.getUserFromQueryParam(ctx, database.GetUserByEmailQuery, email)
}

func (s basicService) GetUserByUsername(ctx context.Context, username string) (user model.User, err error) {
	return s.getUserFromQueryParam(ctx, database.GetUserByUsernameQuery, username)
}

func (s basicService) CreateUser(ctx context.Context, currentuser model.User) (err error) {
	// check for proper input argument
	if unsafe.Sizeof(currentuser) == 0 {
		return helper.ErrNoUserProvided
	}

	err = s.validateUser(err, currentuser)
	if err != nil {
		return err
	}

	// check if user exists already in data store based on
	// id, user name, and email address
	userExists, err := s.database.DoesUserExist(currentuser.Username, "username = ?")

	s.logger.Info("does user exist", zap.Bool("user exists", userExists))

	if userExists == true {
		return helper.ErrUserAlreadyExists
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		s.logger.Error(err.Error())
		return err
	}

	if userExists {
		return helper.ErrUserAlreadyExists
	}

	currentuser, err = s.validateAndHashPassword(currentuser)
	if err != nil {
		return err
	}

	s.logger.Info("Adding User")
	// Create user in database
	err = s.database.CreateUser(currentuser)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}

	// write to the create welcome email queue
	_ = s.ProducerQueues.SendMessageToQueue("Welcome To Lens", "lens_welcome_email")
	return nil
}

// validateAndHashPassword checks if a given user password and confirmed password match
func (s basicService) validateAndHashPassword(currentuser model.User) (user model.User, err error) {
	// check if confirmed password and actual password match
	if currentuser.PassWord != currentuser.PassWordConfirmed {
		s.logger.Error(helper.ErrPasswordsNotEqual.Error())
		return currentuser, helper.ErrPasswordsNotEqual
	}
	//  hash password
	hashedPassword, err := s.hashAndSalt([]byte(currentuser.PassWord))
	if err != nil {
		s.logger.Error(err.Error())
		return currentuser, err
	}
	// reset the hashed passwords
	currentuser.PassWord = hashedPassword
	currentuser.PassWordConfirmed = hashedPassword

	return currentuser, nil
}

func (s basicService) validateUser(err error, currentuser model.User) error {
	// validate fields are present
	err = validate.Struct(currentuser)
	if err != nil {
		// this check is only needed when code could produce
		// an invalid value for validation such as interface with nil value
		if _, ok := err.(*validator.InvalidValidationError); ok {
			s.logger.Error(err.Error())
			return err
		}

		var consolidatedErrMsg = ""
		// translate all error at once
		for _, err := range err.(validator.ValidationErrors) {
			consolidatedErrMsg += "Invalid field " + err.Field() + " "
		}

		s.logger.Error(err.Error())
		return errors.New(consolidatedErrMsg)
	}
	return nil
}

// hashAndSalt hashes and salts a password
func (s basicService) hashAndSalt(pwd []byte) (string, error) {

	// Use GenerateFromPassword to hash & salt pwd
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		s.logger.Error(err.Error())
		return "", err
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash), nil
}

// comparePasswords compares a hashed password and a plain password
func (s basicService) comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		s.logger.Error(err.Error())
		return false
	}
	return true
}

// getUserFromQueryParam obtains a user based on a query parameter
func (s basicService) getUserFromQueryParam(ctx context.Context, query string, param string) (user model.User, err error) {
	user, err = s.database.GetUserBasedOnParam(param, query)
	if err != nil {
		s.logger.Error(err.Error())
		return model.User{}, err
	}

	return user, nil
}
