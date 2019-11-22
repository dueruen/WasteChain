package grpc

import (
	"context"

	kitgrpc "github.com/go-kit/kit/transport/grpc"

	pb "github.com/dueruen/WasteChain/service/blockchain/gen/proto"
	"github.com/dueruen/WasteChain/service/blockchain/pkg/transport"
)

type server struct {
	getShipmentData kitgrpc.Handler
	publish         kitgrpc.Handler
}

func NewGrpcServer(endpoints transport.Endpoints, options []kitgrpc.ServerOption) pb.BlockchainServiceServer {
	return &server{
		getShipmentData: kitgrpc.NewServer(endpoints.GetShipmentData, decodeGetShipmentDataRequest, encodeGetShipmentDataResponse),
		publish:         kitgrpc.NewServer(endpoints.Publish, decodePublishRequest, encodePublishResponse),
	}
}

func (server *server) GetShipmentData(ctx context.Context, req *pb.GetShipmentDataRequest) (*pb.GetShipmentDataResponse, error) {
	_, rep, err := server.getShipmentData.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetShipmentDataResponse), nil
}

func decodeGetShipmentDataRequest(_ context.Context, request interface{}) (interface{}, error) {
	return request.(*pb.GetShipmentDataRequest), nil
}

func encodeGetShipmentDataResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response.(*pb.GetShipmentDataResponse), nil
}

func (server *server) Publish(ctx context.Context, req *pb.PublishRequest) (*pb.PublishResponse, error) {
	_, rep, err := server.publish.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.PublishResponse), nil
}

func decodePublishRequest(_ context.Context, request interface{}) (interface{}, error) {
	return request.(*pb.PublishRequest), nil
}

func encodePublishResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response.(*pb.PublishResponse), nil
}
