package postgres

import (
	"fmt"
	"time"

	pb "github.com/dueruen/WasteChain/service/account/gen/proto"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Storage struct {
	db *gorm.DB
}

func NewStorage(db_string string) (*Storage, error) {
	db := connect(db_string)

	err := createSchema(db)
	if err != nil {
		return nil, err
	}
	return &Storage{db: db}, nil
}

func Close(s *Storage) {
	s.db.Close()
}

func connect(db_string string) *gorm.DB {
	i := 5
	for i > 0 {
		db, err := gorm.Open("postgres", db_string)
		if err != nil {
			fmt.Println("Can't connect to db, sleeping for 2 sec, err: ", err)
			time.Sleep(2 * time.Second)
			i--
			continue
		} else {
			fmt.Println("Connected to storage")
			return db
		}
	}
	panic("Not connected to storage")
}

func createSchema(db *gorm.DB) error {
	if db.HasTable(&pb.Company{}) {
		return nil
	}

	db.AutoMigrate(&pb.Company{}, &pb.Address{}, &pb.ContactInfo{}, &pb.Employee{})

	db.Model(&pb.Company{}).AddForeignKey("address_id", "addresses(id)", "CASCADE", "CASCADE")
	db.Model(&pb.ContactInfo{}).AddForeignKey("company_id", "companies(id)", "CASCADE", "CASCADE")
	db.Model(&pb.Employee{}).AddForeignKey("company_id", "companies(id)", "CASCADE", "CASCADE")
	return nil
}

func (storage *Storage) CreateNewCompany(id string, company *pb.CreateCompany) (*pb.Company, error) {
	newCompany := &pb.Company{
		ID:   id,
		Name: company.Name,
		Address: &pb.Address{
			Number:   company.Address.Number,
			RoadName: company.Address.RoadName,
			ZipCode:  company.Address.ZipCode,
		},
		ContactInfo: []*pb.ContactInfo{
			&pb.ContactInfo{
				Title:       company.ContactInfo.Title,
				Mail:        company.ContactInfo.Mail,
				PhoneNumber: company.ContactInfo.PhoneNumber,
			},
		},
	}

	addressID, _ := uuid.NewV4()
	newCompany.Address.ID = addressID.String()

	contractID, _ := uuid.NewV4()
	newCompany.ContactInfo[0].ID = contractID.String()
	newCompany.ContactInfo[0].CompanyID = newCompany.ID

	storage.db.Create(newCompany)
	return newCompany, nil
}

func (storage *Storage) CreateEmployee(id string, employee *pb.CreateEmployee) (*pb.Employee, error) {
	newEmployee := &pb.Employee{
		ID:        id,
		Name:      employee.Name,
		CompanyID: employee.CompanyID,
	}
	storage.db.Create(newEmployee)
	return newEmployee, nil
}

func (storage *Storage) ListAllCompanies() ([]*pb.Company, error) {
	var companies []*pb.Company
	storage.db.Find(&companies)

	for _, com := range companies {
		com = getAllCompanyData(storage.db, com)
	}
	return companies, nil
}

func (storage *Storage) GetCompany(companyID string) (*pb.Company, error) {
	var company pb.Company
	storage.db.Where("id = ?", companyID).First(&company)
	company = *getAllCompanyData(storage.db, &company)
	return &company, nil
}

func getAllCompanyData(db *gorm.DB, company *pb.Company) *pb.Company {
	var address pb.Address
	db.Where("id = ?", company.AddressID).First(&address)
	company.Address = &address

	var contactInfo []*pb.ContactInfo
	db.Where("company_id = ?", company.ID).Find(&contactInfo)
	company.ContactInfo = contactInfo

	var employees []*pb.Employee
	db.Where("company_id = ?", company.ID).Find(&employees)
	company.Employees = employees
	return company
}

func (storage *Storage) ListAllEmployeesInCompany(companyID string) ([]*pb.Employee, error) {
	var employees []*pb.Employee
	storage.db.Where("company_id = ?", companyID).Find(&employees)
	return employees, nil
}

func (storage *Storage) GetEmployee(employeeID string) (*pb.Employee, error) {
	var employee pb.Employee
	storage.db.Where("id = ?", employeeID).First(&employee)
	return &employee, nil
}
