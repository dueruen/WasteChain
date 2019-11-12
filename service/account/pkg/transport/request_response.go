package transport

import (
	"github.com/dueruen/WasteChain/service/account/pkg/creating"
	"github.com/dueruen/WasteChain/service/account/pkg/listing"
)

type (
	CreateCompanyRequest struct {
		Company *creating.Company
	}

	CreateCompanyResponse struct {
		Error   string
		Company *listing.Company
	}

	ListAllCompaniesRequest struct{}

	ListAllCompaniesResponse struct {
		Error     string
		Companies []*listing.Company
	}

	GetCompanyRequest struct {
		ID string
	}

	GetCompanyResponse struct {
		Error   string
		Company *listing.Company
	}

	CreateEmployeeRequest struct {
		Employee *creating.Employee
	}

	CreateEmployeeResponse struct {
		Error    string
		Employee *listing.Employee
	}

	ListAllEmployeesInCompanyRequest struct {
		ID string
	}

	ListAllEmployeesInCompanyResponse struct {
		Error     string
		Employees []*listing.Employee
	}

	GetEmployeeRequest struct {
		ID string
	}

	GetEmployeeResponse struct {
		Error    string
		Employee *listing.Employee
	}
)
