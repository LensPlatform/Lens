package service

import (
	"context"

	"github.com/go-kit/kit/metrics"
	"go.uber.org/zap"
)

// Service is a CRUD interface definition for the user microservice
//
// CreateUser effectively add a user object to the backend data store
// it takes as input an type interface and returns the object id of the created
// user and any error encountered that may have occurred during this transaction
type Service interface {
	CreateUser(ctx context.Context, user interface{})(id string, err error)
}

// User represents a single user profile
// ID should always be globally unique
type User struct {
	ID string `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	UserName string `json:"user_name"`
	Email string `json:"email"`
	PassWord string `json:"password"`
	Age int `json:"age"`
	BirthDate string `json:"birth_date"`
	PhoneNumber string `json:"phone_number"`
	Addresses []Address `json:"location"`
	Bio string `json:"bio"`
	EducationalExperience []Education `json:"education"`
	UserInterests Interests `json:"interests"`
	Headline string `json:"headline"`
	UserSubscriptions []Subscriptions `json:"subscriptions"`
	Intent string `json:"intent"`
}

type Address struct {
	City string `json:"city"`
	State string `json:"state"`
	Country string `json:"country"`
}

type Education struct{
	InstitutionName []string `json:"institution_name"`
	HighestDegreeEarned string `json:"highest_degree_earned"`
	Graduated bool `json:"graduated"`
	Major string `json:"major"`
	Minor string `json:"minor"`
	YearsOfAttendance string `json:"years_of_attendance"`
}

type Interests struct {
	Industry []Industry `json:"industries_of_interest"`
	Topic []Topic `json:"topics_of_interest"`
}

type Topic struct{
	TopicName string `json:"topic_name"`
	TopicType string `json:"topic_type"`
}

type Industry struct {
	IndustryName string `json:"industry_name"`
}

type Subscriptions struct {
	SubscriptionName string `json:"subscription_name"`
	Subscribe bool `json:"subscribe"`
}

// New returns a basic Service with all of the expected middlewares wired in.
func New(logger *zap.Logger, request, success, failed metrics.Counter) Service {
	var svc Service
	{
		svc = NewBasicService(logger)
		svc = LoggingMiddleware(logger)(svc)
		svc = InstrumentingMiddleware( request, success, failed )(svc)
	}
	return svc
}

// NewBasicService returns a naïve, stateless implementation of Service.
func NewBasicService(logger *zap.Logger) Service {
	return basicService{logger:logger}
}

type basicService struct{
	logger *zap.Logger
}

// CreateUser implements service.
//
// Creates a user in the backend store given some user object of interface type
func (s basicService) CreateUser(ctx context.Context, user interface{}) (id string, err error) {
	if user == nil {
		return "",ErrNoUserProvided
	}
	// Todo: Implement logic for user service
	s.logger.Info("Hello")
	// panic("Implement me")
	return "hello", nil
}

