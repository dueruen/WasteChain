package transport

import (
    "context"
    "github.com/dueruen/WasteChain/service/qr/pkg/creating"
    "github.com/go-kit/kit/endpoint"
    pb "github.com/dueruen/WasteChain/service/qr/gen/proto"
)

type Endpoints struct {
    CreateQRCode endpoint.Endpoint
}


func MakeEndpoints(createService creating.Service) Endpoints { 
    return Endpoints{
        CreateQRCode: makeCreateQRCodeEndpoint(createService)
    }
}


func makeCreateQRCodeEndpoint(service creating.Service) endpoint.Endpoint { 
    return func(ctx context.Context, request interface{}) (interface{}, error) { 
        req := request.(*pb.CreateQRequest)
        res, _ := service.CreateQRCode(req.dataString)
        return &pb.CreateQRResponse{QrCode: res}, nil
    }


}