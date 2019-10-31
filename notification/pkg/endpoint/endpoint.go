package endpoint

import (
	"context"

	endpoint "github.com/go-kit/kit/endpoint"

	service "github.com/LensPlatform/Lens/notification/pkg/service"
)

// SendEmailRequest collects the request parameters for the SendEmail method.
type SendEmailRequest struct {
	Email   string `json:"email"`
	Content string `json:"content"`
}

// SendEmailResponse collects the response parameters for the SendEmail method.
type SendEmailResponse struct {
	Id string `json:"id"`
	E0 error  `json:"e0"`
}

// MakeSendEmailEndpoint returns an endpoint that invokes SendEmail on the service.
func MakeSendEmailEndpoint(s service.NotificationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SendEmailRequest)
		id, e0 := s.SendEmail(ctx, req.Email, req.Content)
		return SendEmailResponse{Id: id, E0: e0}, nil
	}
}

// Failed implements Failer.
func (r SendEmailResponse) Failed() error {
	return r.E0
}

// SendWelcomeToLensEmailRequest collects the request parameters for the SendWelcomeToLensEmail method.
type SendWelcomeToLensEmailRequest struct {
	Email string `json:"email"`
}

// SendWelcomeToLensEmailResponse collects the response parameters for the SendWelcomeToLensEmail method.
type SendWelcomeToLensEmailResponse struct {
	Id string `json:"id"`
	E0 error  `json:"e0"`
}

// MakeSendWelcomeToLensEmailEndpoint returns an endpoint that invokes SendWelcomeToLensEmail on the service.
func MakeSendWelcomeToLensEmailEndpoint(s service.NotificationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SendWelcomeToLensEmailRequest)
		id, e0 := s.SendWelcomeToLensEmail(ctx, req.Email)
		return SendWelcomeToLensEmailResponse{Id: id, E0: e0}, nil
	}
}

// Failed implements Failer.
func (r SendWelcomeToLensEmailResponse) Failed() error {
	return r.E0
}

// SendPasswordChangeEmailRequest collects the request parameters for the SendPasswordChangeEmail method.
type SendPasswordChangeEmailRequest struct {
	Email string `json:"email"`
}

// SendPasswordChangeEmailResponse collects the response parameters for the SendPasswordChangeEmail method.
type SendPasswordChangeEmailResponse struct {
	Id string `json:"id"`
	E0 error  `json:"e0"`
}

// MakeSendPasswordChangeEmailEndpoint returns an endpoint that invokes SendPasswordChangeEmail on the service.
func MakeSendPasswordChangeEmailEndpoint(s service.NotificationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SendPasswordChangeEmailRequest)
		id, e0 := s.SendPasswordChangeEmail(ctx, req.Email)
		return SendPasswordChangeEmailResponse{Id: id, E0: e0}, nil
	}
}

// Failed implements Failer.
func (r SendPasswordChangeEmailResponse) Failed() error {
	return r.E0
}

// SendEmailAccountResetEmailRequest collects the request parameters for the SendEmailAccountResetEmail method.
type SendEmailAccountResetEmailRequest struct {
	OldEmail string `json:"old_email"`
	NewEmail string `json:"new_email"`
}

// SendEmailAccountResetEmailResponse collects the response parameters for the SendEmailAccountResetEmail method.
type SendEmailAccountResetEmailResponse struct {
	Id string `json:"id"`
	E0 error  `json:"e0"`
}

// MakeSendEmailAccountResetEmailEndpoint returns an endpoint that invokes SendEmailAccountResetEmail on the service.
func MakeSendEmailAccountResetEmailEndpoint(s service.NotificationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SendEmailAccountResetEmailRequest)
		id, e0 := s.SendEmailAccountResetEmail(ctx, req.OldEmail, req.NewEmail)
		return SendEmailAccountResetEmailResponse{Id: id, E0: e0}, nil
	}
}

// Failed implements Failer.
func (r SendEmailAccountResetEmailResponse) Failed() error {
	return r.E0
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// SendEmail implements Service. Primarily useful in a client.
func (e Endpoints) SendEmail(ctx context.Context, email string, content string) (e0 error) {
	request := SendEmailRequest{
		Content: content,
		Email:   email,
	}
	response, err := e.SendEmailEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SendEmailResponse).E0
}

// SendWelcomeToLensEmail implements Service. Primarily useful in a client.
func (e Endpoints) SendWelcomeToLensEmail(ctx context.Context, email string) (e0 error) {
	request := SendWelcomeToLensEmailRequest{Email: email}
	response, err := e.SendWelcomeToLensEmailEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SendWelcomeToLensEmailResponse).E0
}

// SendPasswordChangeEmail implements Service. Primarily useful in a client.
func (e Endpoints) SendPasswordChangeEmail(ctx context.Context, email string) (e0 error) {
	request := SendPasswordChangeEmailRequest{Email: email}
	response, err := e.SendPasswordChangeEmailEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SendPasswordChangeEmailResponse).E0
}

// SendEmailAccountResetEmail implements Service. Primarily useful in a client.
func (e Endpoints) SendEmailAccountResetEmail(ctx context.Context, oldEmail string, newEmail string) (e0 error) {
	request := SendEmailAccountResetEmailRequest{
		NewEmail: newEmail,
		OldEmail: oldEmail,
	}
	response, err := e.SendEmailAccountResetEmailEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SendEmailAccountResetEmailResponse).E0
}
