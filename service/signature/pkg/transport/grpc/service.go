package grpc

import (
	"context"

	pb "github.com/dueruen/WasteChain/service/signature/gen/proto"
	"github.com/dueruen/WasteChain/service/signature/pkg/transport"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
)

type server struct {
	singleSign         kitgrpc.Handler
	startDoubleSign    kitgrpc.Handler
	continueDoubleSign kitgrpc.Handler
	singleVerify       kitgrpc.Handler
	doubleVerify       kitgrpc.Handler
	createKeys         kitgrpc.Handler
}

func NewGRPCServer(endpoints transport.Endpoints, options []kitgrpc.ServerOption) pb.SignatureServiceServer {
	return &server{
		singleSign: kitgrpc.NewServer(
			endpoints.SingleSign, decodeSingleSignRequest, encodeSingleSignResponse,
		),
		startDoubleSign: kitgrpc.NewServer(
			endpoints.StartDoubleSign, decodeStartDoubleSignRequest, encodeStartDoubleSignResponse,
		),
		continueDoubleSign: kitgrpc.NewServer(
			endpoints.ContinueDoubleSign, decodeContinueDoubleSignRequest, encodeContinueDoubleSignResponse,
		),
		singleVerify: kitgrpc.NewServer(
			endpoints.SingleVerify, decodeSingleVerifyRequest, encodeSingleVerifyResponse,
		),
		doubleVerify: kitgrpc.NewServer(
			endpoints.DoubleVerify, decodeDoubleVerifyRequest, encodeDoubleVerifyResponse,
		),
		createKeys: kitgrpc.NewServer(
			endpoints.CreateKeys, decodeCreateKeysRequest, encodeCreateKeysResponse,
		),
	}
}

func (srv *server) SingleSign(ctx context.Context, req *pb.SingleSignRequest) (*pb.SingleSignResponse, error) {
	_, rep, err := srv.singleSign.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SingleSignResponse), nil
}

func decodeSingleSignRequest(_ context.Context, request interface{}) (interface{}, error) {
	return request.(*pb.SingleSignRequest), nil
}

func encodeSingleSignResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response.(*pb.SingleSignResponse), nil
}

func (srv *server) StartDoubleSign(ctx context.Context, req *pb.StartDoubleSignRequest) (*pb.StartDoubleSignResponse, error) {
	_, rep, err := srv.startDoubleSign.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.StartDoubleSignResponse), nil
}

func decodeStartDoubleSignRequest(_ context.Context, request interface{}) (interface{}, error) {
	return request.(*pb.StartDoubleSignRequest), nil
}

func encodeStartDoubleSignResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response.(*pb.StartDoubleSignResponse), nil
}

func (srv *server) ContinueDoubleSign(ctx context.Context, req *pb.ContinueDoubleSignRequest) (*pb.ContinueDoubleSignResponse, error) {
	_, rep, err := srv.continueDoubleSign.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ContinueDoubleSignResponse), nil
}

func decodeContinueDoubleSignRequest(_ context.Context, request interface{}) (interface{}, error) {
	return request.(*pb.ContinueDoubleSignRequest), nil
}

func encodeContinueDoubleSignResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response.(*pb.ContinueDoubleSignResponse), nil
}

func (srv *server) SingleVerify(ctx context.Context, req *pb.SingleVerifyRequest) (*pb.SingleVerifyResponse, error) {
	_, rep, err := srv.singleVerify.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SingleVerifyResponse), nil
}

func decodeSingleVerifyRequest(_ context.Context, request interface{}) (interface{}, error) {
	return request.(*pb.SingleVerifyRequest), nil
}

func encodeSingleVerifyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response.(*pb.SingleVerifyResponse), nil
}

func (srv *server) DoubleVerify(ctx context.Context, req *pb.DoubleVerifyRequest) (*pb.DoubleVerifyResponse, error) {
	_, rep, err := srv.doubleVerify.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DoubleVerifyResponse), nil
}

func decodeDoubleVerifyRequest(_ context.Context, request interface{}) (interface{}, error) {
	return request.(*pb.DoubleVerifyRequest), nil
}

func encodeDoubleVerifyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response.(*pb.DoubleVerifyResponse), nil
}

func (srv *server) CreateKeys(ctx context.Context, req *pb.CreateKeysRequest) (*pb.CreateKeysResponse, error) {
	_, rep, err := srv.createKeys.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateKeysResponse), nil
}

func decodeCreateKeysRequest(_ context.Context, request interface{}) (interface{}, error) {
	return request.(*pb.CreateKeysRequest), nil
}

func encodeCreateKeysResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response.(*pb.CreateKeysResponse), nil
}
