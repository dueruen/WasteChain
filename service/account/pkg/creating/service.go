package creating

import (
	"context"
	"errors"

	pb "github.com/dueruen/WasteChain/service/account/gen/proto"
	"github.com/gofrs/uuid"
)

type Service interface {
	CreateCompany(company *pb.CreateCompany) (*pb.Company, error)
	CreateEmployee(employee *pb.CreateEmployee) (*pb.Employee, error)
}

type Repository interface {
	CreateNewCompany(id string, company *pb.CreateCompany) (*pb.Company, error)
	CreateEmployee(id string, employee *pb.CreateEmployee) (*pb.Employee, error)
}

type service struct {
	createRepo Repository
	authClient pb.AuthenticationServiceClient
}

func NewService(createRepo Repository, authClient pb.AuthenticationServiceClient) Service {
	return &service{createRepo, authClient}
}

func (srv *service) CreateCompany(company *pb.CreateCompany) (*pb.Company, error) {
	id, _ := uuid.NewV4()

	res, err := srv.authClient.CreateCredentials(context.Background(), &pb.CreateCredentialsRequest{
		Credentials: &pb.Credentials{
			UserID:   id.String(),
			Username: company.AuthData.Username,
			Password: company.AuthData.Password,
		},
	})
	if err != nil {
		return nil, err
	}
	if res.Error != "" {
		return nil, errors.New(res.Error)
	}

	return srv.createRepo.CreateNewCompany(id.String(), company)
}

func (srv *service) CreateEmployee(employee *pb.CreateEmployee) (*pb.Employee, error) {
	id, _ := uuid.NewV4()

	res, err := srv.authClient.CreateCredentials(context.Background(), &pb.CreateCredentialsRequest{
		Credentials: &pb.Credentials{
			UserID:   id.String(),
			Username: employee.AuthData.Username,
			Password: employee.AuthData.Password,
		},
	})
	if err != nil {
		return nil, err
	}
	if res.Error != "" {
		return nil, errors.New(res.Error)
	}

	return srv.createRepo.CreateEmployee(id.String(), employee)
}
