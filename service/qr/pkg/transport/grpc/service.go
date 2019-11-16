package grpc

import (
	"context"

	kitgrpc "github.com/go-kit/kit/transport/grpc"

	pb "github.com/dueruen/WasteChain/service/qr/gen/proto"
	"github.com/dueruen/WasteChain/service/qr/pkg/transport"
)

type server struct {
	createQRCode kitgrpc.Handler
}

func NewGrpcServer(endpoints transport.Endpoints, options []kitgrpc.ServerOption) pb.QRServiceServer {
	return &server{
		createQRCode: kitgrpc.NewServer(endpoints.CreateQRCode, decodeCreateQRRequest, encodeCreateQRResponse),
	}
}

func (server *server) CreateQRCode(ctx context.Context, req *pb.CreateQRRequest) (*pb.CreateQRResponse, error) {
	_, rep, err := server.createQRCode.ServeGRPC(ctx, req)

	if err != nil {
		return nil, err
	}

	return rep.(*pb.CreateQRResponse), nil
}

func decodeCreateQRRequest(_ context.Context, request interface{}) (interface{}, error) {
	return request.(*pb.CreateQRRequest), nil
}

func encodeCreateQRResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response.(*pb.CreateQRResponse), nil
}
