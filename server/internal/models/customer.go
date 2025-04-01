package models

import "gorm.io/gorm"

// Customer representa um cliente
type Customer struct {
	gorm.Model

	Name                string `gorm:"size:100;not null" json:"name"`
	DocumentType        string `gorm:"size:20" json:"document_type"` // 'cpf', 'cnpj'
	DocumentNumber      string `gorm:"size:20;unique" json:"document_number"`
	Email               string `gorm:"size:100;unique" json:"email"`
	Phone               string `gorm:"size:20" json:"phone"`
	AddressStreet       string `gorm:"size:255" json:"address_street"`
	AddressNumber       string `gorm:"size:20" json:"address_number"`
	AddressComplement   string `gorm:"size:100" json:"address_complement"`
	AddressNeighborhood string `gorm:"size:100" json:"address_neighborhood"`
	AddressCity         string `gorm:"size:100" json:"address_city"`
	AddressState        string `gorm:"size:50" json:"address_state"`
	AddressZipcode      string `gorm:"size:20" json:"address_zipcode"`
	IsActive            bool   `gorm:"default:true" json:"is_active"`
	CreatedByID         *uint  `gorm:"column:created_by" json:"created_by"`
	CreatedBy           *User  `gorm:"foreignKey:CreatedByID" json:"created_by_user,omitempty"`
	Sales               []Sale `gorm:"foreignKey:CustomerID" json:"-"`
}

// TableName especifica o nome da tabela
func (Customer) TableName() string {
	return "customers"
}
