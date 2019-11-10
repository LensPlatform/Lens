package service

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"unsafe"

	"github.com/go-kit/kit/metrics"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
)

// Service is a CRUD interface definition for the user microservice
//
// CreateUser effectively add a user object to the backend data store
// it takes as input an type interface and returns the object id of the created
// user and any error encountered that may have occurred during this transaction
type Service interface {
	CreateUser(ctx context.Context, user User)(err error)
	GetUserById(ctx context.Context, id string)(user User, err error)
	GetUserByEmail(ctx context.Context, email string)(user User, err error)
	GetUserByUsername(ctx context.Context, username string)(user User, err error)
	LogIn(ctx context.Context, username, password string)(user User, err error)
}

// User represents a single user profile
// ID should always be globally unique
type User struct {
	ID string `json:"id" validate:"-" db:"id"`
	FirstName string `json:"first_name" validate:"required" db:"firstname"`
	LastName string `json:"last_name" validate:"required" db:"lastname"`
	UserName string `json:"user_name" validate:"required" db:"username"`
	Gender string `json:"gender" validate:"-" db:"gender"`
	Languages string `json:"Languages" validate:"-" db:"languages"`
	Email string `json:"email" validate:"required,email" db:"email"`
	PassWord string `json:"password" validate:"required,gte=8,lte=20" db:"password"`
	PassWordConfirmed string `json:"password_confirmed" validate:"required,gte=8,lte=20" db:"passwordconf"`
	Age int `json:"age" validate:"gte=0,lte=120" db:"age"`
	BirthDate string `json:"birth_date" validate:"required" db:"birthdate"`
	PhoneNumber string `json:"phone_number" validate:"required" db:"phonenumber"`
	Addresses Address `json:"location" validate:"-" db:"address"`
	Bio string `json:"bio" validate:"required" db:"bio"`
	EducationalExperience Education `json:"education" validate:"-" db:"education"`
	UserInterests Interests `json:"interests" validate:"-" db:"interests"`
	Headline string `json:"headline" validate:"max=30" db:"headline"`
	UserSubscriptions Subscriptions `json:"subscriptions" validate:"-" db:"subscriptions"`
	Intent string `json:"intent" validate:"required" db:"intent"`
	Skills Skillset `json:"skillset" validate:"-" db:"skills"`
}

type JsonEmbeddable struct {}

type Address struct {
	JsonEmbeddable
	City string `json:"city" validate:"required"`
	State string `json:"state" validate:"required"`
	Country string `json:"country" validate:"required"`
}

type Education struct{
	JsonEmbeddable
	MostRecentInstitutionName string `json:"most_recent_institution_name" validate:"required"`
	HighestDegreeEarned string `json:"highest_degree_earned" validate:"required"`
	Graduated bool `json:"graduated" validate:"required"`
	Major string `json:"major" validate:"required"`
	Minor string `json:"minor" validate:"required"`
	YearsOfAttendance string `json:"years_of_attendance" validate:"required"`
}

type Interests struct {
	JsonEmbeddable
	Industry []Industry `json:"industries_of_interest" validate:"omitempty"`
	Topic []Topic `json:"topics_of_interest" validate:"omitempty"`
}

type Topic struct{
	JsonEmbeddable
	TopicName string `json:"topic_name" validate:"required"`
	TopicType string `json:"topic_type" validate:"required"`
}

type Industry struct {
	JsonEmbeddable
	IndustryName string `json:"industry_name" validate:"required"`
}

type Subscriptions struct {
	JsonEmbeddable
	SubscriptionName string `json:"subscription_name" validate:"required"`
	Subscribe bool `json:"subscribe" validate:"required"`
}

type Skillset struct {
	JsonEmbeddable
	Skills []Skill `json:"skills" validate:"required"`
}

type Skill struct {
	JsonEmbeddable
	Type string `json:"skill_type" validate:"required"`
	Name string `json:"skill_name" validate:"required"`
}

var validate = validator.New()

// New returns a basic Service with all of the expected middlewares wired in.
func New(logger *zap.Logger, db *sqlx.DB, CreateUserRequest, successfulCreateUserReq,
	failedCreateUserReq, getUserRequests, successfulGetUserReq, failedGetUserReq, successfulLogInReq, failedLogInReq metrics.Counter) Service {
	var svc Service
	{
		svc = NewBasicService(db, logger)
		svc = LoggingMiddleware(logger)(svc)
		svc = InstrumentingMiddleware(CreateUserRequest, successfulCreateUserReq,
			failedCreateUserReq, getUserRequests, successfulGetUserReq,
			failedGetUserReq,successfulLogInReq, failedLogInReq )(svc)
	}
	return svc
}

// NewBasicService returns a naïve, stateless implementation of Service.
func NewBasicService(db *sqlx.DB, logger *zap.Logger) Service {
	return basicService{logger:logger, dbConn: db}
}

type basicService struct{
	logger *zap.Logger
	dbConn *sqlx.DB
}

func (s basicService) LogIn(ctx context.Context, username, password string) (user User, err error) {
	var userObj User
	if username == ""{
		s.logger.Error(ErrNoUsernameProvided.Error())
		return userObj, ErrNoUsernameProvided
	}

	if password == ""{
		s.logger.Error(ErrNoPasswordProvided.Error())
		return userObj, ErrNoPasswordProvided
	}

	// check if user exists in the database
	err = s.dbConn.QueryRowContext(ctx, GetUserByUsernameQuery, username).Scan(&userObj.ID, &userObj.FirstName, &userObj.LastName,
		&userObj.UserName, &userObj.Email, &userObj.PassWord, &userObj.PassWordConfirmed, &userObj.Age, &userObj.BirthDate, &user.PhoneNumber,
		&userObj.Addresses, &userObj.Bio, &userObj.EducationalExperience, &userObj.UserInterests, &user.Headline, &user.Intent, &user.UserSubscriptions,
		&user.Gender, &user.Languages, &user.Skills)

	if err != nil{
		if err == sql.ErrNoRows{
			s.logger.Error(ErrInvalidUsernameProvided.Error())
			return userObj, ErrInvalidUsernameProvided
		}

		return userObj, err
	}

	s.logger.Info("Password", zap.String("password", userObj.PassWord))

	// check if passwords match
	isEqual := s.comparePasswords(userObj.PassWord, []byte(password))

	if !isEqual{
		s.logger.Error(ErrInvalidPasswordProvided.Error())
		return User{}, ErrInvalidPasswordProvided
	}

	return userObj, nil
}

func (s basicService) GetUserById(ctx context.Context, id string) (user User, err error) {
	return s.getUserFromQueryParam(ctx,GetUserByIdQuery,id)
}

func (s basicService) GetUserByEmail(ctx context.Context, email string) (user User, err error) {
	return s.getUserFromQueryParam(ctx,GetUserByEmailQuery,email)
}

func (s basicService) GetUserByUsername(ctx context.Context, username string) (user User, err error) {
	return s.getUserFromQueryParam(ctx,GetUserByUsernameQuery,username)
}

// CreateUser implements service.
//
// Creates a user in the backend store given some user object of interface type
func (s basicService) CreateUser(ctx context.Context, currentuser User) (err error) {
	// check for proper input argument
	if unsafe.Sizeof(currentuser) == 0 {
		return ErrNoUserProvided
	}

	err = s.validateUser(err, currentuser)
	if err != nil{
		return err
	}

	// check if user exists already in data store based on
	// id, user name, and email address
	rows, err := s.dbConn.QueryContext(ctx, CheckIfUserAlreadyExistQuery,
		currentuser.UserName, currentuser.Email)

	if err != nil {
		s.logger.Error(err.Error())
		return err
	}

	defer rows.Close()

	if rows.Next() != false {
		return ErrUserAlreadyExists
	}

	currentuser, err = s.validateAndHashPassword(currentuser)
	if err != nil {
		return err
	}
	// create a user id
	currentuser.ID = uuid.New().String()

	// Create user in database
	_, err = s.dbConn.Exec(CreateUserQuery, currentuser.FirstName, currentuser.LastName,
						currentuser.UserName, currentuser.Email, currentuser.PassWord,
						currentuser.PassWordConfirmed,currentuser.Age, currentuser.BirthDate,
						currentuser.PhoneNumber, currentuser.Addresses, currentuser.EducationalExperience,
						currentuser.UserInterests,currentuser.Headline,currentuser.Intent,
						currentuser.UserSubscriptions, currentuser.Bio, currentuser.Gender,
						currentuser.Skills, currentuser.Languages)

	if err != nil {
		s.logger.Error(err.Error())
		return err
	}

	return nil
}

func (s basicService) validateAndHashPassword(currentuser User) (user User, err error){
	// check if confirmed password and actual password match
	if currentuser.PassWord != currentuser.PassWordConfirmed {
		s.logger.Error(ErrPasswordsNotEqual.Error())
		return currentuser, ErrPasswordsNotEqual
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

func (s basicService) validateUser(err error, currentuser User) error {
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

// ser struct implements the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
func (u JsonEmbeddable) Value() (driver.Value, error) {
	return json.Marshal(u)
}

// User struct implements the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (u JsonEmbeddable) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &u)
}

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

func (s basicService) getUserFromQueryParam(ctx context.Context, query string, param string) (user User, err error){
	err = s.dbConn.QueryRowContext(ctx, query, param).Scan(&user.ID,&user.FirstName, &user.LastName, &user.UserName, &user.Email,
															&user.PassWord, &user.PassWordConfirmed,
															&user.Age, &user.BirthDate, &user.PhoneNumber,&user.Addresses,
															&user.Bio, &user.EducationalExperience,
															&user.UserInterests, &user.Headline, &user.Intent, &user.UserSubscriptions,
															&user.Gender, &user.Languages,&user.Skills)
	if err != nil {
		s.logger.Error(err.Error())
		return User{}, err
	}
	return user, nil
}


