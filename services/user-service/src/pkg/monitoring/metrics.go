package monitoring

import (
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"

	"github.com/LensPlatform/Lens/services/user-service/src/pkg/service"
)

// InitMetrics Creates the (sparse) metrics used in the service.
func InitMetrics() service.Counters {
	var createUserReq, successfulCreateUserReq, failedCreateUserReq, getUserRequests, successfulGetUserReq,
	failedGetUserReq, successfulLogInReq, failedLogInReq metrics.Counter
	{
		// Business-level metrics.
		createUserReq = prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "users",
			Subsystem: "users",
			Name:      "create_user_requests",
			Help:      "Total count of create user requests via the CreateUser method.",
		}, []string{})
		successfulCreateUserReq = prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "users",
			Subsystem: "users",
			Name:      "create_user_success_ops",
			Help:      "Total count of successful create user requests via the CreateUser method.",
		}, []string{})
		failedCreateUserReq = prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "users",
			Subsystem: "users",
			Name:      "create_user_failed_ops",
			Help:      "Total count of failed create user requests via the CreateUser method.",
		}, []string{})
		getUserRequests = prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "users",
			Subsystem: "users",
			Name:      "get_user_requests",
			Help:      "Total count of get user requests.",
		}, []string{})
		successfulGetUserReq = prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "users",
			Subsystem: "users",
			Name:      "get_user_requests_success_ops",
			Help:      "Total count of successful get user requests.",
		}, []string{})
		failedGetUserReq = prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "users",
			Subsystem: "users",
			Name:      "get_user_requests_failed_ops",
			Help:      "Total count of failed get user requests.",
		}, []string{})
		successfulLogInReq = prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "users",
			Subsystem: "users",
			Name:      "login_requests_sucess_ops",
			Help:      "Total count of successful logIn requests.",
		}, []string{})
		failedLogInReq = prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "users",
			Subsystem: "users",
			Name:      "login_requests_failed_ops",
			Help:      "Total count of failed login requests.",
		}, []string{})
	}

	var duration metrics.Histogram
	{
		// Endpoint-level metrics.
		duration = prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "users",
			Subsystem: "users",
			Name:      "request_duration_seconds",
			Help:      "Request duration in seconds.",
		}, []string{"method", "success"})
	}

	counter := service.Counters{
		CreateUserRequest:           createUserReq,
		SuccessfulCreateUserRequest: successfulCreateUserReq,
		FailedCreateUserRequest:     failedCreateUserReq,
		GetUserRequest:              getUserRequests,
		SuccessfulGetUserRequest:    successfulGetUserReq,
		FailedGetUserRequest:        failedGetUserReq,
		SuccessfulLogInRequest:      successfulLogInReq,
		FailedLogInRequest:          failedLogInReq,
		Duration:                    duration,
	}

	return counter
}