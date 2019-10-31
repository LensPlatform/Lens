package grpc

import (
	"context"
	"errors"

	endpoint "github.com/LensPlatform/Lens/social/pkg/endpoint"
	pb "github.com/LensPlatform/Lens/social/pkg/grpc/pb"
	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
)

// makeFollowUserHandler creates the handler logic
func makeFollowUserHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.FollowUserEndpoint, decodeFollowUserRequest, encodeFollowUserResponse, options...)
}

// decodeFollowUserResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain FollowUser request.
// TODO implement the decoder
func decodeFollowUserRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Social' Decoder is not impelemented")
}

// encodeFollowUserResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeFollowUserResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Social' Encoder is not impelemented")
}
func (g *grpcServer) FollowUser(ctx context1.Context, req *pb.FollowUserRequest) (*pb.FollowUserReply, error) {
	_, rep, err := g.followUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.FollowUserReply), nil
}

// makeUnFollowerUserHandler creates the handler logic
func makeUnFollowerUserHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UnFollowerUserEndpoint, decodeUnFollowerUserRequest, encodeUnFollowerUserResponse, options...)
}

// decodeUnFollowerUserResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain UnFollowerUser request.
// TODO implement the decoder
func decodeUnFollowerUserRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Social' Decoder is not impelemented")
}

// encodeUnFollowerUserResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeUnFollowerUserResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Social' Encoder is not impelemented")
}
func (g *grpcServer) UnFollowerUser(ctx context1.Context, req *pb.UnFollowerUserRequest) (*pb.UnFollowerUserReply, error) {
	_, rep, err := g.unFollowerUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UnFollowerUserReply), nil
}

// makeFollowGroupHandler creates the handler logic
func makeFollowGroupHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.FollowGroupEndpoint, decodeFollowGroupRequest, encodeFollowGroupResponse, options...)
}

// decodeFollowGroupResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain FollowGroup request.
// TODO implement the decoder
func decodeFollowGroupRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Social' Decoder is not impelemented")
}

// encodeFollowGroupResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeFollowGroupResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Social' Encoder is not impelemented")
}
func (g *grpcServer) FollowGroup(ctx context1.Context, req *pb.FollowGroupRequest) (*pb.FollowGroupReply, error) {
	_, rep, err := g.followGroup.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.FollowGroupReply), nil
}

// makeUnFollowGroupHandler creates the handler logic
func makeUnFollowGroupHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UnFollowGroupEndpoint, decodeUnFollowGroupRequest, encodeUnFollowGroupResponse, options...)
}

// decodeUnFollowGroupResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain UnFollowGroup request.
// TODO implement the decoder
func decodeUnFollowGroupRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Social' Decoder is not impelemented")
}

// encodeUnFollowGroupResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeUnFollowGroupResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Social' Encoder is not impelemented")
}
func (g *grpcServer) UnFollowGroup(ctx context1.Context, req *pb.UnFollowGroupRequest) (*pb.UnFollowGroupReply, error) {
	_, rep, err := g.unFollowGroup.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UnFollowGroupReply), nil
}

// makeFollowTeamHandler creates the handler logic
func makeFollowTeamHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.FollowTeamEndpoint, decodeFollowTeamRequest, encodeFollowTeamResponse, options...)
}

// decodeFollowTeamResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain FollowTeam request.
// TODO implement the decoder
func decodeFollowTeamRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Social' Decoder is not impelemented")
}

// encodeFollowTeamResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeFollowTeamResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Social' Encoder is not impelemented")
}
func (g *grpcServer) FollowTeam(ctx context1.Context, req *pb.FollowTeamRequest) (*pb.FollowTeamReply, error) {
	_, rep, err := g.followTeam.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.FollowTeamReply), nil
}

// makeUnFollowTeamHandler creates the handler logic
func makeUnFollowTeamHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UnFollowTeamEndpoint, decodeUnFollowTeamRequest, encodeUnFollowTeamResponse, options...)
}

// decodeUnFollowTeamResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain UnFollowTeam request.
// TODO implement the decoder
func decodeUnFollowTeamRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Social' Decoder is not impelemented")
}

// encodeUnFollowTeamResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeUnFollowTeamResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Social' Encoder is not impelemented")
}
func (g *grpcServer) UnFollowTeam(ctx context1.Context, req *pb.UnFollowTeamRequest) (*pb.UnFollowTeamReply, error) {
	_, rep, err := g.unFollowTeam.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UnFollowTeamReply), nil
}
