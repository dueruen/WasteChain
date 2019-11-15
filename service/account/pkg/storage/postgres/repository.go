package postgres

import (
	"fmt"

	"github.com/dueruen/WasteChain/service/account/pkg/creating"
	"github.com/dueruen/WasteChain/service/account/pkg/listing"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Storage struct {
	db *gorm.DB
}

func NewStorage(host, port, user, dbname, password string) (*Storage, error) {
	db, err := connect(host, port, user, dbname, password)
	if err != nil {

		return nil, err
	}
	err = createSchema(db)
	if err != nil {
		return nil, err
	}
	return &Storage{db: db}, nil
}

func Close(s *Storage) {
	s.db.Close()
}

func connect(host, port, user, dbname, password string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host="+host+" port="+port+" user="+user+" dbname="+dbname+" password="+password+" sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func createSchema(db *gorm.DB) error {
	if db.HasTable(&Company{}) {
		return nil
	}

	db.AutoMigrate(&Company{}, &Address{}, &ContactInfo{}, &Employee{})

	db.Model(&Company{}).AddForeignKey("address_id", "addresses(id)", "CASCADE", "CASCADE")
	db.Model(&ContactInfo{}).AddForeignKey("company_id", "companies(id)", "CASCADE", "CASCADE")
	db.Model(&Employee{}).AddForeignKey("company_id", "companies(id)", "CASCADE", "CASCADE")
	return nil
}

func (storage *Storage) CreateNewCompany(company *creating.Company) (*listing.Company, error) {
	newCompany := mapCreateCompany(company)
	id, _ := uuid.NewV4()
	newCompany.ID = id.String()

	addressID, _ := uuid.NewV4()
	newCompany.Address.ID = addressID.String()

	for _, contact := range newCompany.ContactInfo {
		contractID, _ := uuid.NewV4()
		contact.ID = contractID.String()
		contact.CompanyID = newCompany.ID
	}
	storage.db.Create(newCompany)
	return mapToListCompany(newCompany), nil
}

func (storage *Storage) CreateEmployee(employee *creating.Employee) (*listing.Employee, error) {
	fmt.Println("create emp: ", employee)
	newEmployee := mapCreateEmployee(employee)
	uuid, _ := uuid.NewV4()
	newEmployee.ID = uuid.String()
	fmt.Println("New emp: ", newEmployee)
	storage.db.Create(newEmployee)
	return mapToListEmployee(newEmployee), nil
}

func (storage *Storage) ListAllCompanies() ([]*listing.Company, error) {
	var companies []*Company
	storage.db.Find(&companies)

	for _, com := range companies {
		com = getAllCompanyData(storage.db, com)
	}
	for _, com := range companies {
		fmt.Printf("All: %v\n", com)
	}

	return mapAllToListCompany(companies), nil
}

func (storage *Storage) GetCompany(companyID string) (*listing.Company, error) {
	var company Company
	storage.db.Where("id = ?", companyID).First(&company)
	company = *getAllCompanyData(storage.db, &company)
	return mapToListCompany(&company), nil
}

func getAllCompanyData(db *gorm.DB, company *Company) *Company {
	var address Address
	db.Where("id = ?", company.AddressID).First(&address)
	company.Address = &address

	var contactInfo []*ContactInfo
	db.Where("company_id = ?", company.ID).Find(&contactInfo)
	company.ContactInfo = contactInfo

	var employees []*Employee
	db.Where("company_id = ?", company.ID).Find(&employees)
	company.Employees = employees
	return company
}

func (storage *Storage) ListAllEmployeesInCompany(companyID string) ([]*listing.Employee, error) {
	var employees []*Employee
	storage.db.Where("company_id = ?", companyID).Find(&employees)
	return mapAllToListEmployee(employees), nil
}

func (storage *Storage) GetEmployee(employeeID string) (*listing.Employee, error) {
	var employee Employee
	storage.db.Where("id = ?", employeeID).First(&employee)
	return mapToListEmployee(&employee), nil
}
