package models

import "gorm.io/gorm"

// PaymentMethod representa um método de pagamento
type PaymentMethod struct {
	gorm.Model

	Name        string `gorm:"size:50;not null;unique" json:"name"`
	Description string `json:"description"`
	IsActive    bool   `gorm:"default:true" json:"is_active"`
}

// TableName especifica o nome da tabela
func (PaymentMethod) TableName() string {
	return "payment_methods"
}
