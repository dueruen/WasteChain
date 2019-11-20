package grpc

import (
	"context"

	kitgrpc "github.com/go-kit/kit/transport/grpc"

	pb "github.com/dueruen/WasteChain/service/shipment/gen/proto"
	"github.com/dueruen/WasteChain/service/shipment/pkg/transport"
)

type server struct {
	createShipment     kitgrpc.Handler
	transferShipment   kitgrpc.Handler
	processShipment    kitgrpc.Handler
	listAllShipments   kitgrpc.Handler
	getShipmentDetails kitgrpc.Handler
}

func NewGrpcServer(endpoints transport.Endpoints, options []kitgrpc.ServerOption) pb.ShipmentServiceServer {
	return &server{
		createShipment:     kitgrpc.NewServer(endpoints.CreateShipment, decodeCreateShipmentRequest, encodeCreateShipmentResponse),
		transferShipment:   kitgrpc.NewServer(endpoints.TransferShipment, decodeTransferShipmentRequest, encodeTransferShipmentResponse),
		processShipment:    kitgrpc.NewServer(endpoints.ProcessShipment, decodeProcessShipmentRequest, encodeProcessShipmentResponse),
		listAllShipments:   kitgrpc.NewServer(endpoints.ListAllShipments, decodeListAllShipmentsRequest, encodeListAllShipmentsResponse),
		getShipmentDetails: kitgrpc.NewServer(endpoints.GetShipmentDetails, decodeGetShipmentDetailsRequest, encodeGetShipmentDetailsResponse),
	}
}

func (server *server) CreateShipment(ctx context.Context, req *pb.CreateShipmentRequest) (*pb.CreateShipmentResponse, error) {
	_, rep, err := server.createShipment.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateShipmentResponse), nil
}

func decodeCreateShipmentRequest(_ context.Context, request interface{}) (interface{}, error) {
	return request.(*pb.CreateShipmentRequest), nil
}

func encodeCreateShipmentResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response.(*pb.CreateShipmentResponse), nil
}

func (server *server) TransferShipment(ctx context.Context, req *pb.TransferShipmentRequest) (*pb.TransferShipmentResponse, error) {
	_, rep, err := server.transferShipment.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.TransferShipmentResponse), nil
}

func decodeTransferShipmentRequest(_ context.Context, request interface{}) (interface{}, error) {
	return request.(*pb.TransferShipmentRequest), nil
}

func encodeTransferShipmentResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response.(*pb.TransferShipmentResponse), nil
}

func (server *server) ProcessShipment(ctx context.Context, req *pb.ProcessShipmentRequest) (*pb.ProcessShipmentResponse, error) {
	_, rep, err := server.processShipment.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ProcessShipmentResponse), nil
}

func decodeProcessShipmentRequest(_ context.Context, request interface{}) (interface{}, error) {
	return request.(*pb.ProcessShipmentRequest), nil
}

func encodeProcessShipmentResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response.(*pb.ProcessShipmentResponse), nil
}

func (server *server) ListAllShipments(ctx context.Context, req *pb.ListAllShipmentsRequest) (*pb.ListAllShipmentsResponse, error) {
	_, rep, err := server.listAllShipments.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.ListAllShipmentsResponse), nil
}

func decodeListAllShipmentsRequest(_ context.Context, request interface{}) (interface{}, error) {
	return request.(*pb.ListAllShipmentsRequest), nil
}

func encodeListAllShipmentsResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response.(*pb.ListAllShipmentsResponse), nil
}

func (server *server) GetShipmentDetails(ctx context.Context, req *pb.GetShipmentDetailsRequest) (*pb.GetShipmentDetailsResponse, error) {
	_, rep, err := server.getShipmentDetails.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetShipmentDetailsResponse), nil
}

func decodeGetShipmentDetailsRequest(_ context.Context, request interface{}) (interface{}, error) {
	return request.(*pb.GetShipmentDetailsRequest), nil
}

func encodeGetShipmentDetailsResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response.(*pb.GetShipmentDetailsResponse), nil
}
