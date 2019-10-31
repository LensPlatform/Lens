package endpoint

import (
	"context"

	service "LensPlatform/Lens/notifications/pkg/service"

	endpoint "github.com/go-kit/kit/endpoint"
)

// SendEmailRequest collects the request parameters for the SendEmail method.
type SendEmailRequest struct {
	Email   string `json:"email"`
	Content string `json:"content"`
}

// SendEmailResponse collects the response parameters for the SendEmail method.
type SendEmailResponse struct {
	S0 string `json:"s0"`
	E1 error  `json:"e1"`
}

// MakeSendEmailEndpoint returns an endpoint that invokes SendEmail on the service.
func MakeSendEmailEndpoint(s service.NotificationsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SendEmailRequest)
		s0, e1 := s.SendEmail(ctx, req.Email, req.Content)
		return SendEmailResponse{
			E1: e1,
			S0: s0,
		}, nil
	}
}

// Failed implements Failer.
func (r SendEmailResponse) Failed() error {
	return r.E1
}

// SendWelcomeToLensEmailRequest collects the request parameters for the SendWelcomeToLensEmail method.
type SendWelcomeToLensEmailRequest struct {
	Email string `json:"email"`
}

// SendWelcomeToLensEmailResponse collects the response parameters for the SendWelcomeToLensEmail method.
type SendWelcomeToLensEmailResponse struct {
	S0 string `json:"s0"`
	E1 error  `json:"e1"`
}

// MakeSendWelcomeToLensEmailEndpoint returns an endpoint that invokes SendWelcomeToLensEmail on the service.
func MakeSendWelcomeToLensEmailEndpoint(s service.NotificationsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SendWelcomeToLensEmailRequest)
		s0, e1 := s.SendWelcomeToLensEmail(ctx, req.Email)
		return SendWelcomeToLensEmailResponse{
			E1: e1,
			S0: s0,
		}, nil
	}
}

// Failed implements Failer.
func (r SendWelcomeToLensEmailResponse) Failed() error {
	return r.E1
}

// SendPasswordChangeEmailRequest collects the request parameters for the SendPasswordChangeEmail method.
type SendPasswordChangeEmailRequest struct {
	Email string `json:"email"`
}

// SendPasswordChangeEmailResponse collects the response parameters for the SendPasswordChangeEmail method.
type SendPasswordChangeEmailResponse struct {
	S0 string `json:"s0"`
	E1 error  `json:"e1"`
}

// MakeSendPasswordChangeEmailEndpoint returns an endpoint that invokes SendPasswordChangeEmail on the service.
func MakeSendPasswordChangeEmailEndpoint(s service.NotificationsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SendPasswordChangeEmailRequest)
		s0, e1 := s.SendPasswordChangeEmail(ctx, req.Email)
		return SendPasswordChangeEmailResponse{
			E1: e1,
			S0: s0,
		}, nil
	}
}

// Failed implements Failer.
func (r SendPasswordChangeEmailResponse) Failed() error {
	return r.E1
}

// SendEmailAccountResetEmailRequest collects the request parameters for the SendEmailAccountResetEmail method.
type SendEmailAccountResetEmailRequest struct {
	OldEmail string `json:"old_email"`
	NewEmail string `json:"new_email"`
}

// SendEmailAccountResetEmailResponse collects the response parameters for the SendEmailAccountResetEmail method.
type SendEmailAccountResetEmailResponse struct {
	S0 string `json:"s0"`
	E1 error  `json:"e1"`
}

// MakeSendEmailAccountResetEmailEndpoint returns an endpoint that invokes SendEmailAccountResetEmail on the service.
func MakeSendEmailAccountResetEmailEndpoint(s service.NotificationsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SendEmailAccountResetEmailRequest)
		s0, e1 := s.SendEmailAccountResetEmail(ctx, req.OldEmail, req.NewEmail)
		return SendEmailAccountResetEmailResponse{
			E1: e1,
			S0: s0,
		}, nil
	}
}

// Failed implements Failer.
func (r SendEmailAccountResetEmailResponse) Failed() error {
	return r.E1
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// SendEmail implements Service. Primarily useful in a client.
func (e Endpoints) SendEmail(ctx context.Context, email string, content string) (s0 string, e1 error) {
	request := SendEmailRequest{
		Content: content,
		Email:   email,
	}
	response, err := e.SendEmailEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SendEmailResponse).S0, response.(SendEmailResponse).E1
}

// SendWelcomeToLensEmail implements Service. Primarily useful in a client.
func (e Endpoints) SendWelcomeToLensEmail(ctx context.Context, email string) (s0 string, e1 error) {
	request := SendWelcomeToLensEmailRequest{Email: email}
	response, err := e.SendWelcomeToLensEmailEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SendWelcomeToLensEmailResponse).S0, response.(SendWelcomeToLensEmailResponse).E1
}

// SendPasswordChangeEmail implements Service. Primarily useful in a client.
func (e Endpoints) SendPasswordChangeEmail(ctx context.Context, email string) (s0 string, e1 error) {
	request := SendPasswordChangeEmailRequest{Email: email}
	response, err := e.SendPasswordChangeEmailEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SendPasswordChangeEmailResponse).S0, response.(SendPasswordChangeEmailResponse).E1
}

// SendEmailAccountResetEmail implements Service. Primarily useful in a client.
func (e Endpoints) SendEmailAccountResetEmail(ctx context.Context, oldEmail string, newEmail string) (s0 string, e1 error) {
	request := SendEmailAccountResetEmailRequest{
		NewEmail: newEmail,
		OldEmail: oldEmail,
	}
	response, err := e.SendEmailAccountResetEmailEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SendEmailAccountResetEmailResponse).S0, response.(SendEmailAccountResetEmailResponse).E1
}
