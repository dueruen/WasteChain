package transport

import (
	"context"

	pb "github.com/dueruen/WasteChain/service/shipment/gen/proto"
	"github.com/dueruen/WasteChain/service/shipment/pkg/creating"
	"github.com/dueruen/WasteChain/service/shipment/pkg/listing"
	"github.com/dueruen/WasteChain/service/shipment/pkg/processing"
	"github.com/dueruen/WasteChain/service/shipment/pkg/transfering"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateShipment     endpoint.Endpoint
	TransferShipment   endpoint.Endpoint
	ProcessShipment    endpoint.Endpoint
	GetShipmentDetails endpoint.Endpoint
	ListAllShipments   endpoint.Endpoint
}

func MakeEndpoints(createService creating.Service, listingService listing.Service,
	transferService transfering.Service, processingService processing.Service) Endpoints {
	return Endpoints{
		CreateShipment:     makeCreateShipmentEndpoint(createService),
		TransferShipment:   makeTransferShipmentEndpoint(transferService),
		ProcessShipment:    makeProcessShipmentEndpoint(processingService),
		GetShipmentDetails: makeGetShipmentDetailsEndpoint(listingService),
		ListAllShipments:   makeListAllShipmentsEndpoint(listingService),
	}
}

func makeCreateShipmentEndpoint(service creating.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.CreateShipmentRequest)
		res, _ := service.CreateShipment(req)
		return &pb.CreateShipmentResponse{ID: res}, nil
	}
}

func makeTransferShipmentEndpoint(service transfering.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.TransferShipmentRequest)
		res, err := service.TransferShipment(req)
		if err != nil {
			return &pb.TransferShipmentResponse{Error: err.Error()}, err
		}
		return res, nil
	}
}

func makeProcessShipmentEndpoint(service processing.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.ProcessShipmentRequest)
		service.ProcessShipment(req)
		return &pb.ProcessShipmentResponse{}, nil
	}
}

func makeGetShipmentDetailsEndpoint(service listing.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.GetShipmentDetailsRequest)
		_, res := service.GetShipmentDetails(req)
		return &pb.GetShipmentDetailsResponse{Shipment: res}, nil
	}
}

func makeListAllShipmentsEndpoint(service listing.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_, res := service.ListAllShipments()
		return &pb.ListAllShipmentsResponse{ShipmentList: res}, nil
	}
}
