package transport

import (
	"context"

	pb "github.com/dueruen/WasteChain/service/blockchain/gen/proto"
	"github.com/dueruen/WasteChain/service/blockchain/pkg/publish"
	"github.com/dueruen/WasteChain/service/blockchain/pkg/receive"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetShipmentData endpoint.Endpoint
	Publish         endpoint.Endpoint
}

func MakeEndpoints(receiveSrv receive.Service, publishSrv publish.Service) Endpoints {
	return Endpoints{
		GetShipmentData: makeGetShipmentDataEndpoint(receiveSrv),
		Publish:         makePublishEndpoint(publishSrv),
	}
}

func makeGetShipmentDataEndpoint(service receive.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.GetShipmentDataRequest)
		err, history := service.GetShipmentData(req.ShipmentID)
		if err != nil {
			return &pb.GetShipmentDataResponse{Error: err.Error()}, nil
		}
		return &pb.GetShipmentDataResponse{History: history}, nil
	}
}

func makePublishEndpoint(service publish.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.PublishRequest)
		res := service.Publish(req.ShipmentID, req.Data)
		if res != nil {
			return &pb.PublishResponse{Error: res.Error()}, nil
		}
		return &pb.PublishResponse{}, nil
	}
}
