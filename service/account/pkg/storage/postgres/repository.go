package postgres

import (
	"context"

	pb "github.com/dueruen/WasteChain/service/account/gen/proto"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/gofrs/uuid"
)

type Storage struct {
	db *pg.DB
}

func NewStorage(user, password, database string) (*Storage, error) {
	db, err := connect(user, password, database)
	if err != nil {
		return nil, err
	}
	return &Storage{db: db}, nil
}

func connect(user, password, database string) (*pg.DB, error) {
	db := pg.Connect(&pg.Options{
		User:     user,
		Password: password,
		Database: database,
	})
	defer db.Close()

	err := createSchema(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*Employee)(nil), (*Company)(nil), (*Address)(nil), (*ContactInfo)(nil)} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (storage *Storage) CreateNewCompany(company *pb.CreateCompanyRequest) (*pb.CreateCompanyResponse, error) {
	uuid, _ := uuid.NewV4()
	newCompany := Company{
		ID:   uuid.String(),
		Name: company.Name,
		Address: &Address{
			RoadName: company.Address.RoadName,
			Number:   company.Address.Number,
			ZipCode:  company.Address.ZipCode,
		},
		ContactInfo: &ContactInfo{
			Mail:        company.ContactInfo.Mail,
			PhoneNumber: company.ContactInfo.PhoneNumber,
		},
	}
	err := storage.db.Insert(newCompany)
	if err != nil {
		return nil, err
	}
	return &pb.CreateCompanyResponse{
		Company: &pb.Company{
			ID:   newCompany.ID,
			Name: newCompany.Name,
			Address: &pb.Address{
				RoadName: newCompany.Address.RoadName,
				Number:   newCompany.Address.Number,
				ZipCode:  newCompany.Address.ZipCode,
			},
			ContactInfo: &pb.ContactInfo{
				Mail:        newCompany.ContactInfo.Mail,
				PhoneNumber: newCompany.ContactInfo.PhoneNumber,
			},
		},
	}, nil
}

func (storage *Storage) CreateEmployee(*pb.CreateEmployeeRequest) (*pb.CreateEmployeeResponse, error) {

}

func (storage *Storage) ListAllCompanies(ctx context.Context, in *pb.ListAllCompaniesRequest) (*pb.ListAllCompaniesResponse, error) {

}

func (storage *Storage) GetCompany(ctx context.Context, in *pb.GetCompanyRequest) (*pb.GetCompanyResponse, error) {

}

func (storage *Storage) ListAllEmployeesInCompany(ctx context.Context, in *pb.ListAllEmployeesInCompanyRequest) (*pb.ListAllEmployeesInCompanyResponse, error) {

}

func (storage *Storage) GetEmployee(ctx context.Context, in *pb.GetEmployeeRequest) (*pb.GetEmployeeResponse, error) {

}
