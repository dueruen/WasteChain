package postgres

type Company struct {
	ID          string `gorm:"primary_key"`
	Name        string
	Address     *Address
	AddressID   string         `gorm:"ForeignKey:id"`
	ContactInfo []*ContactInfo `gorm:"ForeignKey:CompanyID"`
	Employees   []*Employee    `gorm:"ForeignKey:CompanyID"`
}

type Address struct {
	ID       string `gorm:"primary_key"`
	RoadName string
	Number   int32
	ZipCode  int32
}

type ContactInfo struct {
	ID          string `gorm:"primary_key"`
	Title       string
	PhoneNumber int32
	Mail        string
	CompanyID   string
}

type Employee struct {
	ID        string `gorm:"primary_key"`
	Name      string
	CompanyID string
}
