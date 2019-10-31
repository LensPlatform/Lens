package grpc

import (
	"context"
	"errors"
	endpoint "github.com/LensPlatform/Lens/users/pkg/endpoint"
	pb "github.com/LensPlatform/Lens/users/pkg/grpc/pb"
	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
)

// makeCreateUserAccountHandler creates the handler logic
func makeCreateUserAccountHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreateUserAccountEndpoint, decodeCreateUserAccountRequest, encodeCreateUserAccountResponse, options...)
}

// decodeCreateUserAccountResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain CreateUserAccount request.
// TODO implement the decoder
func decodeCreateUserAccountRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Decoder is not impelemented")
}

// encodeCreateUserAccountResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeCreateUserAccountResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Encoder is not impelemented")
}
func (g *grpcServer) CreateUserAccount(ctx context1.Context, req *pb.CreateUserAccountRequest) (*pb.CreateUserAccountReply, error) {
	_, rep, err := g.createUserAccount.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateUserAccountReply), nil
}

// makeUpdateUserAccountHandler creates the handler logic
func makeUpdateUserAccountHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UpdateUserAccountEndpoint, decodeUpdateUserAccountRequest, encodeUpdateUserAccountResponse, options...)
}

// decodeUpdateUserAccountResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain UpdateUserAccount request.
// TODO implement the decoder
func decodeUpdateUserAccountRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Decoder is not impelemented")
}

// encodeUpdateUserAccountResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeUpdateUserAccountResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Encoder is not impelemented")
}
func (g *grpcServer) UpdateUserAccount(ctx context1.Context, req *pb.UpdateUserAccountRequest) (*pb.UpdateUserAccountReply, error) {
	_, rep, err := g.updateUserAccount.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdateUserAccountReply), nil
}

// makeDeleteUserAccountHandler creates the handler logic
func makeDeleteUserAccountHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.DeleteUserAccountEndpoint, decodeDeleteUserAccountRequest, encodeDeleteUserAccountResponse, options...)
}

// decodeDeleteUserAccountResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain DeleteUserAccount request.
// TODO implement the decoder
func decodeDeleteUserAccountRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Decoder is not impelemented")
}

// encodeDeleteUserAccountResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeDeleteUserAccountResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Encoder is not impelemented")
}
func (g *grpcServer) DeleteUserAccount(ctx context1.Context, req *pb.DeleteUserAccountRequest) (*pb.DeleteUserAccountReply, error) {
	_, rep, err := g.deleteUserAccount.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteUserAccountReply), nil
}

// makeGetUserAccountHandler creates the handler logic
func makeGetUserAccountHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetUserAccountEndpoint, decodeGetUserAccountRequest, encodeGetUserAccountResponse, options...)
}

// decodeGetUserAccountResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetUserAccount request.
// TODO implement the decoder
func decodeGetUserAccountRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Decoder is not impelemented")
}

// encodeGetUserAccountResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetUserAccountResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Encoder is not impelemented")
}
func (g *grpcServer) GetUserAccount(ctx context1.Context, req *pb.GetUserAccountRequest) (*pb.GetUserAccountReply, error) {
	_, rep, err := g.getUserAccount.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetUserAccountReply), nil
}

// makeCreateTeamsAccountHandler creates the handler logic
func makeCreateTeamsAccountHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreateTeamsAccountEndpoint, decodeCreateTeamsAccountRequest, encodeCreateTeamsAccountResponse, options...)
}

// decodeCreateTeamsAccountResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain CreateTeamsAccount request.
// TODO implement the decoder
func decodeCreateTeamsAccountRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Decoder is not impelemented")
}

// encodeCreateTeamsAccountResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeCreateTeamsAccountResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Encoder is not impelemented")
}
func (g *grpcServer) CreateTeamsAccount(ctx context1.Context, req *pb.CreateTeamsAccountRequest) (*pb.CreateTeamsAccountReply, error) {
	_, rep, err := g.createTeamsAccount.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateTeamsAccountReply), nil
}

// makeUpdateTeamsAccountHandler creates the handler logic
func makeUpdateTeamsAccountHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UpdateTeamsAccountEndpoint, decodeUpdateTeamsAccountRequest, encodeUpdateTeamsAccountResponse, options...)
}

// decodeUpdateTeamsAccountResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain UpdateTeamsAccount request.
// TODO implement the decoder
func decodeUpdateTeamsAccountRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Decoder is not impelemented")
}

// encodeUpdateTeamsAccountResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeUpdateTeamsAccountResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Encoder is not impelemented")
}
func (g *grpcServer) UpdateTeamsAccount(ctx context1.Context, req *pb.UpdateTeamsAccountRequest) (*pb.UpdateTeamsAccountReply, error) {
	_, rep, err := g.updateTeamsAccount.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdateTeamsAccountReply), nil
}

// makeDeleteTeamsAccountHandler creates the handler logic
func makeDeleteTeamsAccountHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.DeleteTeamsAccountEndpoint, decodeDeleteTeamsAccountRequest, encodeDeleteTeamsAccountResponse, options...)
}

// decodeDeleteTeamsAccountResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain DeleteTeamsAccount request.
// TODO implement the decoder
func decodeDeleteTeamsAccountRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Decoder is not impelemented")
}

// encodeDeleteTeamsAccountResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeDeleteTeamsAccountResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Encoder is not impelemented")
}
func (g *grpcServer) DeleteTeamsAccount(ctx context1.Context, req *pb.DeleteTeamsAccountRequest) (*pb.DeleteTeamsAccountReply, error) {
	_, rep, err := g.deleteTeamsAccount.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteTeamsAccountReply), nil
}

// makeGetTeamsAccountHandler creates the handler logic
func makeGetTeamsAccountHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetTeamsAccountEndpoint, decodeGetTeamsAccountRequest, encodeGetTeamsAccountResponse, options...)
}

// decodeGetTeamsAccountResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetTeamsAccount request.
// TODO implement the decoder
func decodeGetTeamsAccountRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Decoder is not impelemented")
}

// encodeGetTeamsAccountResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetTeamsAccountResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Encoder is not impelemented")
}
func (g *grpcServer) GetTeamsAccount(ctx context1.Context, req *pb.GetTeamsAccountRequest) (*pb.GetTeamsAccountReply, error) {
	_, rep, err := g.getTeamsAccount.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetTeamsAccountReply), nil
}

// makeAddUserToTeamsAccountHandler creates the handler logic
func makeAddUserToTeamsAccountHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.AddUserToTeamsAccountEndpoint, decodeAddUserToTeamsAccountRequest, encodeAddUserToTeamsAccountResponse, options...)
}

// decodeAddUserToTeamsAccountResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain AddUserToTeamsAccount request.
// TODO implement the decoder
func decodeAddUserToTeamsAccountRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Decoder is not impelemented")
}

// encodeAddUserToTeamsAccountResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeAddUserToTeamsAccountResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Encoder is not impelemented")
}
func (g *grpcServer) AddUserToTeamsAccount(ctx context1.Context, req *pb.AddUserToTeamsAccountRequest) (*pb.AddUserToTeamsAccountReply, error) {
	_, rep, err := g.addUserToTeamsAccount.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AddUserToTeamsAccountReply), nil
}

// makeDeleteUserFromTeamsAccountHandler creates the handler logic
func makeDeleteUserFromTeamsAccountHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.DeleteUserFromTeamsAccountEndpoint, decodeDeleteUserFromTeamsAccountRequest, encodeDeleteUserFromTeamsAccountResponse, options...)
}

// decodeDeleteUserFromTeamsAccountResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain DeleteUserFromTeamsAccount request.
// TODO implement the decoder
func decodeDeleteUserFromTeamsAccountRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Decoder is not impelemented")
}

// encodeDeleteUserFromTeamsAccountResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeDeleteUserFromTeamsAccountResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Encoder is not impelemented")
}
func (g *grpcServer) DeleteUserFromTeamsAccount(ctx context1.Context, req *pb.DeleteUserFromTeamsAccountRequest) (*pb.DeleteUserFromTeamsAccountReply, error) {
	_, rep, err := g.deleteUserFromTeamsAccount.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteUserFromTeamsAccountReply), nil
}

// makeCreateGroupHandler creates the handler logic
func makeCreateGroupHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreateGroupEndpoint, decodeCreateGroupRequest, encodeCreateGroupResponse, options...)
}

// decodeCreateGroupResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain CreateGroup request.
// TODO implement the decoder
func decodeCreateGroupRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Decoder is not impelemented")
}

// encodeCreateGroupResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeCreateGroupResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Encoder is not impelemented")
}
func (g *grpcServer) CreateGroup(ctx context1.Context, req *pb.CreateGroupRequest) (*pb.CreateGroupReply, error) {
	_, rep, err := g.createGroup.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateGroupReply), nil
}

// makeSubscribeToGroupHandler creates the handler logic
func makeSubscribeToGroupHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.SubscribeToGroupEndpoint, decodeSubscribeToGroupRequest, encodeSubscribeToGroupResponse, options...)
}

// decodeSubscribeToGroupResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain SubscribeToGroup request.
// TODO implement the decoder
func decodeSubscribeToGroupRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Decoder is not impelemented")
}

// encodeSubscribeToGroupResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeSubscribeToGroupResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Encoder is not impelemented")
}
func (g *grpcServer) SubscribeToGroup(ctx context1.Context, req *pb.SubscribeToGroupRequest) (*pb.SubscribeToGroupReply, error) {
	_, rep, err := g.subscribeToGroup.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SubscribeToGroupReply), nil
}

// makeUnsubscribeFromGroupHandler creates the handler logic
func makeUnsubscribeFromGroupHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UnsubscribeFromGroupEndpoint, decodeUnsubscribeFromGroupRequest, encodeUnsubscribeFromGroupResponse, options...)
}

// decodeUnsubscribeFromGroupResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain UnsubscribeFromGroup request.
// TODO implement the decoder
func decodeUnsubscribeFromGroupRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Decoder is not impelemented")
}

// encodeUnsubscribeFromGroupResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeUnsubscribeFromGroupResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Encoder is not impelemented")
}
func (g *grpcServer) UnsubscribeFromGroup(ctx context1.Context, req *pb.UnsubscribeFromGroupRequest) (*pb.UnsubscribeFromGroupReply, error) {
	_, rep, err := g.unsubscribeFromGroup.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UnsubscribeFromGroupReply), nil
}

// makeDeleteGroupHandler creates the handler logic
func makeDeleteGroupHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.DeleteGroupEndpoint, decodeDeleteGroupRequest, encodeDeleteGroupResponse, options...)
}

// decodeDeleteGroupResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain DeleteGroup request.
// TODO implement the decoder
func decodeDeleteGroupRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Decoder is not impelemented")
}

// encodeDeleteGroupResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeDeleteGroupResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Encoder is not impelemented")
}
func (g *grpcServer) DeleteGroup(ctx context1.Context, req *pb.DeleteGroupRequest) (*pb.DeleteGroupReply, error) {
	_, rep, err := g.deleteGroup.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteGroupReply), nil
}

// makeUpdateGroupHandler creates the handler logic
func makeUpdateGroupHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UpdateGroupEndpoint, decodeUpdateGroupRequest, encodeUpdateGroupResponse, options...)
}

// decodeUpdateGroupResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain UpdateGroup request.
// TODO implement the decoder
func decodeUpdateGroupRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Decoder is not impelemented")
}

// encodeUpdateGroupResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeUpdateGroupResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Encoder is not impelemented")
}
func (g *grpcServer) UpdateGroup(ctx context1.Context, req *pb.UpdateGroupRequest) (*pb.UpdateGroupReply, error) {
	_, rep, err := g.updateGroup.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdateGroupReply), nil
}

// makeGetGroupByIDHandler creates the handler logic
func makeGetGroupByIDHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetGroupByIDEndpoint, decodeGetGroupByIDRequest, encodeGetGroupByIDResponse, options...)
}

// decodeGetGroupByIDResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetGroupByID request.
// TODO implement the decoder
func decodeGetGroupByIDRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Decoder is not impelemented")
}

// encodeGetGroupByIDResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetGroupByIDResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Encoder is not impelemented")
}
func (g *grpcServer) GetGroupByID(ctx context1.Context, req *pb.GetGroupByIDRequest) (*pb.GetGroupByIDReply, error) {
	_, rep, err := g.getGroupByID.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetGroupByIDReply), nil
}

// makeGetGroupByNameHandler creates the handler logic
func makeGetGroupByNameHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetGroupByNameEndpoint, decodeGetGroupByNameRequest, encodeGetGroupByNameResponse, options...)
}

// decodeGetGroupByNameResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetGroupByName request.
// TODO implement the decoder
func decodeGetGroupByNameRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Decoder is not impelemented")
}

// encodeGetGroupByNameResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetGroupByNameResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Encoder is not impelemented")
}
func (g *grpcServer) GetGroupByName(ctx context1.Context, req *pb.GetGroupByNameRequest) (*pb.GetGroupByNameReply, error) {
	_, rep, err := g.getGroupByName.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetGroupByNameReply), nil
}

// makeIsGroupPrivateHandler creates the handler logic
func makeIsGroupPrivateHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.IsGroupPrivateEndpoint, decodeIsGroupPrivateRequest, encodeIsGroupPrivateResponse, options...)
}

// decodeIsGroupPrivateResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain IsGroupPrivate request.
// TODO implement the decoder
func decodeIsGroupPrivateRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Decoder is not impelemented")
}

// encodeIsGroupPrivateResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeIsGroupPrivateResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Users' Encoder is not impelemented")
}
func (g *grpcServer) IsGroupPrivate(ctx context1.Context, req *pb.IsGroupPrivateRequest) (*pb.IsGroupPrivateReply, error) {
	_, rep, err := g.isGroupPrivate.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.IsGroupPrivateReply), nil
}
