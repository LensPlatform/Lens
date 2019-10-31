package grpc

import (
	"context"
	_ "errors"

	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"

	endpoint "github.com/LensPlatform/Lens/notification/pkg/endpoint"
	pb "github.com/LensPlatform/Lens/notification/pkg/grpc/pb"
)

// makeSendEmailHandler creates the handler logic
func makeSendEmailHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.SendEmailEndpoint, decodeSendEmailRequest, encodeSendEmailResponse, options...)
}

// decodeSendEmailResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain SendEmail request.
func decodeSendEmailRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.SendEmailRequest)
	return endpoint.SendEmailRequest{Email: req.Email, Content: req.Content} ,nil
}

// encodeSendEmailResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeSendEmailResponse(_ context.Context, r interface{}) (interface{}, error) {
	reply := r.(endpoint.SendEmailResponse)
	return &pb.SendEmailReply{Id: reply.Id},nil
}
func (g *grpcServer) SendEmail(ctx context1.Context, req *pb.SendEmailRequest) (*pb.SendEmailReply, error) {
	_, rep, err := g.sendEmail.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SendEmailReply), nil
}

// makeSendWelcomeToLensEmailHandler creates the handler logic
func makeSendWelcomeToLensEmailHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.SendWelcomeToLensEmailEndpoint, decodeSendWelcomeToLensEmailRequest, encodeSendWelcomeToLensEmailResponse, options...)
}

// decodeSendWelcomeToLensEmailResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain SendWelcomeToLensEmail request.
func decodeSendWelcomeToLensEmailRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.SendWelcomeToLensEmailRequest)
	return endpoint.SendWelcomeToLensEmailRequest{Email:req.Email},nil
}

// encodeSendWelcomeToLensEmailResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeSendWelcomeToLensEmailResponse(_ context.Context, r interface{}) (interface{}, error) {
	reply := r.(endpoint.SendWelcomeToLensEmailResponse)
	return &pb.SendWelcomeToLensEmailReply{Id: reply.Id},nil
}
func (g *grpcServer) SendWelcomeToLensEmail(ctx context1.Context, req *pb.SendWelcomeToLensEmailRequest) (*pb.SendWelcomeToLensEmailReply, error) {
	_, rep, err := g.sendWelcomeToLensEmail.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SendWelcomeToLensEmailReply), nil
}

// makeSendPasswordChangeEmailHandler creates the handler logic
func makeSendPasswordChangeEmailHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.SendPasswordChangeEmailEndpoint, decodeSendPasswordChangeEmailRequest, encodeSendPasswordChangeEmailResponse, options...)
}

// decodeSendPasswordChangeEmailResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain SendPasswordChangeEmail request.
func decodeSendPasswordChangeEmailRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.SendPasswordChangeEmailRequest)
	return endpoint.SendPasswordChangeEmailRequest{Email:req.Email},nil
}

// encodeSendPasswordChangeEmailResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeSendPasswordChangeEmailResponse(_ context.Context, r interface{}) (interface{}, error) {
	reply := r.(endpoint.SendPasswordChangeEmailResponse)
	return &pb.SendPasswordChangeEmailReply{Id: reply.Id},nil
}
func (g *grpcServer) SendPasswordChangeEmail(ctx context1.Context, req *pb.SendPasswordChangeEmailRequest) (*pb.SendPasswordChangeEmailReply, error) {
	_, rep, err := g.sendPasswordChangeEmail.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SendPasswordChangeEmailReply), nil
}

// makeSendEmailAccountResetEmailHandler creates the handler logic
func makeSendEmailAccountResetEmailHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.SendEmailAccountResetEmailEndpoint, decodeSendEmailAccountResetEmailRequest, encodeSendEmailAccountResetEmailResponse, options...)
}

// decodeSendEmailAccountResetEmailResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain SendEmailAccountResetEmail request.
func decodeSendEmailAccountResetEmailRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.SendEmailAccountResetEmailRequest)
	return endpoint.SendEmailAccountResetEmailRequest{OldEmail: req.OldEmail, NewEmail: req.NewEmail},nil
}

// encodeSendEmailAccountResetEmailResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeSendEmailAccountResetEmailResponse(_ context.Context, r interface{}) (interface{}, error) {
	reply := r.(endpoint.SendEmailAccountResetEmailResponse)
	return &pb.SendEmailAccountResetEmailReply{Id: reply.Id},nil
}
func (g *grpcServer) SendEmailAccountResetEmail(ctx context1.Context, req *pb.SendEmailAccountResetEmailRequest) (*pb.SendEmailAccountResetEmailReply, error) {
	_, rep, err := g.sendEmailAccountResetEmail.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SendEmailAccountResetEmailReply), nil
}
