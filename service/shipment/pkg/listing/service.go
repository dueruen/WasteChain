package listing

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/dueruen/WasteChain/service/shipment/gen/proto"
	encode "github.com/dueruen/WasteChain/service/shipment/pkg/encode"
)

type Service interface {
	GetShipmentDetails(*pb.GetShipmentDetailsRequest) (error, *pb.Shipment)
	ListAllShipments() (error, []*pb.Shipment)
}

type Repository interface {
	GetShipmentDetails(*pb.GetShipmentDetailsRequest) (error, *pb.Shipment)
	ListAllShipments() (error, []*pb.Shipment)
}

type service struct {
	listRepo   Repository
	signClient pb.SignatureServiceClient
}

func NewService(listRepo Repository, signClient pb.SignatureServiceClient) Service {
	return &service{listRepo, signClient}
}

func (srv *service) GetShipmentDetails(request *pb.GetShipmentDetailsRequest) (error, *pb.Shipment) {
	err, data := srv.listRepo.GetShipmentDetails(request)
	if err != nil {
		return err, nil
	}

	res, err := srv.check(data)

	if err != nil {
		fmt.Println("ERR val: ", err)
		return err, &pb.Shipment{}
	}
	if res.Ok {
		return nil, data
	}

	fmt.Println("NOT VALID")
	return errors.New("Data not valid"), &pb.Shipment{}
}

func (srv *service) ListAllShipments() (error, []*pb.Shipment) {
	err, shipments := srv.listRepo.ListAllShipments()
	out := make([]*pb.Shipment, 0)
	if err != nil {
		return err, out
	}

	for _, shipment := range shipments {
		res, err := srv.check(shipment)

		if err != nil {
			fmt.Println("ERR val: ", err)
			//TODO handle error
			out = append(out, &pb.Shipment{})
		}
		if res.Ok {
			out = append(out, shipment)
		}
	}
	return nil, out
}

func (srv *service) check(data *pb.Shipment) (*pb.VerifyHistoryResponse, error) {
	if len(data.History) == 0 {
		return nil, errors.New("No history")
	}

	hist := make([]*pb.VerifyHistoryItemData, 0)
	for _, item := range data.History {
		hist = append(hist, &pb.VerifyHistoryItemData{
			CurrentHolderID: item.OwnerID,
			NewHolderID:     item.ReceiverID,
			Data:            encode.ToByte(item),
		})
	}

	res, err := srv.signClient.VerifyHistory(context.Background(), &pb.VerifyHistoryRequest{
		ShipmentID: data.ID,
		History:    hist,
	})
	return res, err
}
