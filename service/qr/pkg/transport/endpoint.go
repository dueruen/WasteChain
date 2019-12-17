package transport

import (
	"context"

	pb "github.com/dueruen/WasteChain/service/qr/gen/proto"
	"github.com/dueruen/WasteChain/service/qr/pkg/creating"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateQRCode endpoint.Endpoint
}

func MakeEndpoints(createService creating.Service) Endpoints {
	return Endpoints{
		CreateQRCode: makeCreateQRCodeEndpoint(createService),
	}
}

func makeCreateQRCodeEndpoint(service creating.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.CreateQRRequest)
		qr, err := service.CreateQRCode(req.ID, req.DataString)
		if err != nil {
			return &pb.CreateQRResponse{Error: err.Error()}, err
		}
		return &pb.CreateQRResponse{QRCode: *qr}, nil
	}

}
