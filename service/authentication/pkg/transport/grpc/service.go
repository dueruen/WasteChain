package grpc

import (
	"context"

	kitgrpc "github.com/go-kit/kit/transport/grpc"

	pb "github.com/dueruen/WasteChain/service/authentication/gen/proto"
	"github.com/dueruen/WasteChain/service/authentication/pkg/transport"
)

type server struct {
	createCredentials kitgrpc.Handler
	login             kitgrpc.Handler
	validate          kitgrpc.Handler
}

func NewGrpcServer(endpoints transport.Endpoints, options []kitgrpc.ServerOption) pb.AuthenticationServiceServer {
	return &server{
		createCredentials: kitgrpc.NewServer(endpoints.CreateCredentials, decodeCreateCredentialsRequest, encodeCreateCredentialsResponse),
		login:             kitgrpc.NewServer(endpoints.Login, decodeLoginRequest, encodeLoginResponse),
		validate:          kitgrpc.NewServer(endpoints.Validate, decodeValidateRequest, encodeValidateResponse),
	}
}

func (server *server) CreateCredentials(ctx context.Context, req *pb.CreateCredentialsRequest) (*pb.CreateCredentialsResponse, error) {
	_, rep, err := server.createCredentials.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateCredentialsResponse), nil
}

func decodeCreateCredentialsRequest(_ context.Context, request interface{}) (interface{}, error) {
	return request.(*pb.CreateCredentialsRequest), nil
}

func encodeCreateCredentialsResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response.(*pb.CreateCredentialsResponse), nil
}

func (server *server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	_, rep, err := server.login.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.LoginResponse), nil
}

func decodeLoginRequest(_ context.Context, request interface{}) (interface{}, error) {
	return request.(*pb.LoginRequest), nil
}

func encodeLoginResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response.(*pb.LoginResponse), nil
}

func (server *server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	_, rep, err := server.validate.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ValidateResponse), nil
}

func decodeValidateRequest(_ context.Context, request interface{}) (interface{}, error) {
	return request.(*pb.ValidateRequest), nil
}

func encodeValidateResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response.(*pb.ValidateResponse), nil
}
