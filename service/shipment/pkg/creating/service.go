package creating

import (
	"context"
	"fmt"
	"time"

	pb "github.com/dueruen/WasteChain/service/shipment/gen/proto"
	encode "github.com/dueruen/WasteChain/service/shipment/pkg/encode"
)

type Service interface {
	CreateShipment(creationRequest *pb.CreateShipmentRequest) (string, error)
}

type Repository interface {
	CreateNewShipment(creationRequest *pb.CreateShipmentRequest, timestamp string, companyID string) (string, *pb.HistoryItem, error)
}

type service struct {
	createRepo    Repository
	signClient    pb.SignatureServiceClient
	accountClient pb.AccountServiceClient
}

func NewService(createRepo Repository, signClient pb.SignatureServiceClient, accountClient pb.AccountServiceClient) Service {
	return &service{createRepo, signClient, accountClient}
}

func (srv *service) CreateShipment(creationRequest *pb.CreateShipmentRequest) (string, error) {
	fmt.Println("CREATE REQ: ", creationRequest.String())
	res, err := srv.accountClient.GetEmployee(context.Background(), &pb.GetEmployeeRequest{
		ID: creationRequest.CurrentHolderID,
	})
	if err != nil {
		fmt.Println("CREATE ERR: ", err)
	}
	fmt.Println("CREATE: ", res)
	employee := res.Employee

	id, historyItem, err := srv.createRepo.CreateNewShipment(creationRequest, time.Now().String(), employee.CompanyID)
	byteEvent := encode.ToByte(historyItem)

	srv.signClient.SingleSign(context.Background(), &pb.SingleSignRequest{
		Data:       byteEvent,
		UserID:     creationRequest.CurrentHolderID,
		Password:   creationRequest.Password,
		ShipmentID: id,
	})
	fmt.Println("CREATE: ", err)
	return id, err
}
