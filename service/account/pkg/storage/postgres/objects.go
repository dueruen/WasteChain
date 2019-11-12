package postgres

import "time"

type Company struct {
	ID          string `gorm:"primary_key"`
	Name        string
	Address     *Address
	AddressID   string         `gorm:"foreignkey:id"`
	ContactInfo []*ContactInfo `gorm:"foreignkey:CompanyID"`
	Employees   []*Employee    `gorm:"foreignkey:CompanyID"`
	CreatedAt   time.Time
}

type Address struct {
	ID        string `gorm:"primary_key"`
	RoadName  string
	Number    int32
	ZipCode   int32
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ContactInfo struct {
	ID          string `gorm:"primary_key"`
	Title       string
	PhoneNumber int32
	Mail        string
	CompanyID   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Employee struct {
	ID        string `gorm:"primary_key"`
	Name      string
	CompanyID string
	CreatedAt time.Time
	UpdatedAt time.Time
}
