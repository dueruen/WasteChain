package creating

import (
	"github.com/dueruen/WasteChain/service/account/pkg/listing"
)

type Service interface {
	CreateCompany(company *Company) (*listing.Company, error)
	CreateEmployee(employee *Employee) (*listing.Employee, error)
}

type Repository interface {
	CreateNewCompany(company *Company) (*listing.Company, error)
	CreateEmployee(employee *Employee) (*listing.Employee, error)
}

type service struct {
	createRepo Repository
}

func NewService(createRepo Repository) Service {
	return &service{createRepo}
}

func (srv *service) CreateCompany(company *Company) (*listing.Company, error) {
	// TODO make validation
	return srv.createRepo.CreateNewCompany(company)
}

func (srv *service) CreateEmployee(employee *Employee) (*listing.Employee, error) {
	// TODO make validation
	return srv.createRepo.CreateEmployee(employee)
}
