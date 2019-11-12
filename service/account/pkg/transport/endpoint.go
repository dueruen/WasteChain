package transport

import (
	"context"

	"github.com/dueruen/WasteChain/service/account/pkg/creating"
	"github.com/dueruen/WasteChain/service/account/pkg/listing"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateCompany             endpoint.Endpoint
	CreateEmployee            endpoint.Endpoint
	ListAllCompanies          endpoint.Endpoint
	GetCompany                endpoint.Endpoint
	ListAllEmployeesInCompany endpoint.Endpoint
	GetEmployee               endpoint.Endpoint
}

func MakeEndpoints(createService creating.Service, listService listing.Service) Endpoints {
	return Endpoints{
		CreateCompany:             makeCreateCompanyEndpoint(createService),
		CreateEmployee:            makeCreateEmployeeEndpoint(createService),
		ListAllCompanies:          makeListAllCompaniesEndpoint(listService),
		GetCompany:                makeGetCompanyEndpoint(listService),
		ListAllEmployeesInCompany: makeListAllEmployeesInCompanyEndpoint(listService),
		GetEmployee:               makeGetEmployeeEndpoint(listService),
	}
}

func makeCreateCompanyEndpoint(service creating.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateCompanyRequest)
		res, _ := service.CreateCompany(req.Company)
		return CreateCompanyResponse{Company: res}, nil
	}
}

func makeCreateEmployeeEndpoint(service creating.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateEmployeeRequest)
		res, _ := service.CreateEmployee(req.Employee)
		return CreateEmployeeResponse{Employee: res}, nil
	}
}

func makeListAllCompaniesEndpoint(service listing.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//req := request.(ListAllCompaniesRequest)
		res, _ := service.ListAllCompanies()
		return ListAllCompaniesResponse{Companies: res}, nil
	}
}

func makeGetCompanyEndpoint(service listing.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetCompanyRequest)
		res, _ := service.GetCompany(req.ID)
		return GetCompanyResponse{Company: res}, nil
	}
}

func makeListAllEmployeesInCompanyEndpoint(service listing.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ListAllEmployeesInCompanyRequest)
		res, _ := service.ListAllEmployeesInCompany(req.ID)
		return ListAllEmployeesInCompanyResponse{Employees: res}, nil
	}
}

func makeGetEmployeeEndpoint(service listing.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetEmployeeRequest)
		res, _ := service.GetEmployee(req.ID)
		return GetEmployeeResponse{Employee: res}, nil
	}
}
