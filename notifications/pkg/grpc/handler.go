package grpc

import (
	"context"
	"errors"

	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"

	endpoint "LensPlatform/Lens/notifications/pkg/endpoint"
	pb "LensPlatform/Lens/notifications/pkg/grpc/pb"
)

// makeSendEmailHandler creates the handler logic
func makeSendEmailHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.SendEmailEndpoint, decodeSendEmailRequest, encodeSendEmailResponse, options...)
}

// decodeSendEmailResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain SendEmail request.
// TODO implement the decoder
func decodeSendEmailRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Notifications' Decoder is not impelemented")
}

// encodeSendEmailResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeSendEmailResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Notifications' Encoder is not impelemented")
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
// TODO implement the decoder
func decodeSendWelcomeToLensEmailRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Notifications' Decoder is not impelemented")
}

// encodeSendWelcomeToLensEmailResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeSendWelcomeToLensEmailResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Notifications' Encoder is not impelemented")
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
// TODO implement the decoder
func decodeSendPasswordChangeEmailRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Notifications' Decoder is not impelemented")
}

// encodeSendPasswordChangeEmailResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeSendPasswordChangeEmailResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Notifications' Encoder is not impelemented")
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
// TODO implement the decoder
func decodeSendEmailAccountResetEmailRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Notifications' Decoder is not impelemented")
}

// encodeSendEmailAccountResetEmailResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeSendEmailAccountResetEmailResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Notifications' Encoder is not impelemented")
}
func (g *grpcServer) SendEmailAccountResetEmail(ctx context1.Context, req *pb.SendEmailAccountResetEmailRequest) (*pb.SendEmailAccountResetEmailReply, error) {
	_, rep, err := g.sendEmailAccountResetEmail.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SendEmailAccountResetEmailReply), nil
}
