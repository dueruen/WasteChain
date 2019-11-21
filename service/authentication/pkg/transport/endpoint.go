package transport

import (
	"context"

	pb "github.com/dueruen/WasteChain/service/authentication/gen/proto"
	"github.com/dueruen/WasteChain/service/authentication/pkg/auth"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateCredentials endpoint.Endpoint
	Login             endpoint.Endpoint
	Validate          endpoint.Endpoint
}

func MakeEndpoints(srv auth.Service) Endpoints {
	return Endpoints{
		CreateCredentials: makeCreateCredentialsEndpoint(srv),
		Login:             makeLoginEndpoint(srv),
		Validate:          makeValidateEndpoint(srv),
	}
}

func makeCreateCredentialsEndpoint(service auth.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.CreateCredentialsRequest)
		res := service.CreateCredentials(req)
		return res, nil
	}
}

func makeLoginEndpoint(service auth.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.LoginRequest)
		res := service.Login(req)
		return res, nil
	}
}

func makeValidateEndpoint(service auth.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.ValidateRequest)
		res := service.Validate(req)
		return res, nil
	}
}
