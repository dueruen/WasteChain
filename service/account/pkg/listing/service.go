package listing

import (
	pb "github.com/dueruen/WasteChain/service/account/gen/proto"
)

type Service interface {
	ListAllCompanies() ([]*pb.Company, error)
	GetCompany(companyID string) (*pb.Company, error)
	ListAllEmployeesInCompany(companyID string) ([]*pb.Employee, error)
	GetEmployee(employeeID string) (*pb.Employee, error)
}

type Repository interface {
	ListAllCompanies() ([]*pb.Company, error)
	GetCompany(companyID string) (*pb.Company, error)
	ListAllEmployeesInCompany(companyID string) ([]*pb.Employee, error)
	GetEmployee(employeeID string) (*pb.Employee, error)
}

type service struct {
	listRepo Repository
}

func NewService(listRepo Repository) Service {
	return &service{listRepo}
}

func (srv *service) ListAllCompanies() ([]*pb.Company, error) {
	return srv.listRepo.ListAllCompanies()
}

func (srv *service) GetCompany(companyID string) (*pb.Company, error) {
	return srv.listRepo.GetCompany(companyID)
}

func (srv *service) ListAllEmployeesInCompany(companyID string) ([]*pb.Employee, error) {
	return srv.listRepo.ListAllEmployeesInCompany(companyID)
}

func (srv *service) GetEmployee(employeeID string) (*pb.Employee, error) {
	return srv.listRepo.GetEmployee(employeeID)
}
