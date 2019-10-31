package endpoint

import (
	service "LensPlatform/Lens/notifications/pkg/service"
	"context"

	endpoint "github.com/go-kit/kit/endpoint"
	context1 "golang.org/x/net/context"
)

// SendEmailRequest collects the request parameters for the SendEmail method.
type SendEmailRequest struct {
	Email   string `json:"email"`
	Content string `json:"content"`
}

// SendEmailResponse collects the response parameters for the SendEmail method.
type SendEmailResponse struct {
	Id  string `json:"id"`
	Err error  `json:"err"`
}

// MakeSendEmailEndpoint returns an endpoint that invokes SendEmail on the service.
func MakeSendEmailEndpoint(s service.NotificationsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SendEmailRequest)
		id, err := s.SendEmail(ctx, req.Email, req.Content)
		return SendEmailResponse{
			Err: err,
			Id:  id,
		}, nil
	}
}

// Failed implements Failer.
func (r SendEmailResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// SendEmail implements Service. Primarily useful in a client.
func (e Endpoints) SendEmail(ctx context1.Context, email string, content string) (id string, err error) {
	request := SendEmailRequest{
		Content: content,
		Email:   email,
	}
	response, err := e.SendEmailEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SendEmailResponse).Id, response.(SendEmailResponse).Err
}
