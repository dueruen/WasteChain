package transport

import (
	"context"

	pb "github.com/dueruen/WasteChain/service/signature/gen/proto"
	"github.com/dueruen/WasteChain/service/signature/pkg/key"
	"github.com/dueruen/WasteChain/service/signature/pkg/sign"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	SingleSign         endpoint.Endpoint
	StartDoubleSign    endpoint.Endpoint
	ContinueDoubleSign endpoint.Endpoint
	SingleVerify       endpoint.Endpoint
	DoubleVerify       endpoint.Endpoint
	CreateKeys         endpoint.Endpoint
}

func MakeEndpoints(keySrv key.Service, signSrv sign.Service) Endpoints {
	return Endpoints{
		SingleSign:         makeSingleSignEndpoint(signSrv),
		StartDoubleSign:    makeStartDoubleSignEndpoint(signSrv),
		ContinueDoubleSign: makeContinueDoubleSignEndpoint(signSrv),
		SingleVerify:       makeSingleVerifyEndpoint(signSrv),
		DoubleVerify:       makeDoubleVerifyEndpoint(signSrv),
		CreateKeys:         makeCreateKeysEndpoint(keySrv),
	}
}

func makeSingleSignEndpoint(srv sign.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.SingleSignRequest)
		err := srv.SingleSign(req)
		if err != nil {
			return &pb.SingleSignResponse{
				Error: err.Error(),
			}, err
		}
		return &pb.SingleSignResponse{}, nil
	}
}

func makeStartDoubleSignEndpoint(srv sign.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.StartDoubleSignRequest)
		err := srv.StartDoubleSign(req)
		if err != nil {
			return &pb.StartDoubleSignResponse{
				Error: err.Error(),
			}, err
		}
		return &pb.StartDoubleSignResponse{}, nil
	}
}

func makeContinueDoubleSignEndpoint(srv sign.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.ContinueDoubleSignRequest)
		err := srv.ContinueDoubleSign(req)
		if err != nil {
			return &pb.ContinueDoubleSignResponse{
				Error: err.Error(),
			}, err
		}
		return &pb.ContinueDoubleSignResponse{}, nil
	}
}

func makeSingleVerifyEndpoint(srv sign.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.SingleVerifyRequest)
		ok, err := srv.SingleVerify(req)
		if err != nil {
			return &pb.SingleVerifyResponse{
				Ok:    ok,
				Error: err.Error(),
			}, err
		}
		return &pb.SingleVerifyResponse{Ok: ok}, nil
	}
}

func makeDoubleVerifyEndpoint(srv sign.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.DoubleVerifyRequest)
		ok, err := srv.DoubleVerify(req)
		if err != nil {
			return &pb.DoubleVerifyResponse{
				Ok:    ok,
				Error: err.Error(),
			}, err
		}
		return &pb.DoubleVerifyResponse{Ok: ok}, nil
	}
}

func makeCreateKeysEndpoint(srv key.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.CreateKeysRequest)
		err := srv.CreateKeyPair(req.UserID, req.Passphrase)
		if err != nil {
			return &pb.CreateKeysResponse{
				Error: err.Error(),
			}, err
		}
		return &pb.CreateKeysResponse{}, nil
	}
}
