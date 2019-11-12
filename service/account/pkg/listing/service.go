package listing

type Service interface {
	ListAllCompanies() ([]*Company, error)
	GetCompany(companyID string) (*Company, error)
	ListAllEmployeesInCompany(companyID string) ([]*Employee, error)
	GetEmployee(employeeID string) (*Employee, error)
}

type Repository interface {
	ListAllCompanies() ([]*Company, error)
	GetCompany(companyID string) (*Company, error)
	ListAllEmployeesInCompany(companyID string) ([]*Employee, error)
	GetEmployee(employeeID string) (*Employee, error)
}

type service struct {
	listRepo Repository
}

func NewService(listRepo Repository) Service {
	return &service{listRepo}
}

func (srv *service) ListAllCompanies() ([]*Company, error) {
	return srv.listRepo.ListAllCompanies()
}

func (srv *service) GetCompany(companyID string) (*Company, error) {
	return srv.listRepo.GetCompany(companyID)
}

func (srv *service) ListAllEmployeesInCompany(companyID string) ([]*Employee, error) {
	return srv.listRepo.ListAllEmployeesInCompany(companyID)
}

func (srv *service) GetEmployee(employeeID string) (*Employee, error) {
	return srv.listRepo.GetEmployee(employeeID)
}
