package models

import (
	"time"

	"gorm.io/gorm"
)

// Sale representa uma venda
type Sale struct {
	gorm.Model

	Code            string         `gorm:"size:20;not null;unique" json:"code"`
	CustomerID      *uint          `json:"customer_id"`
	Customer        *Customer      `gorm:"foreignKey:CustomerID" json:"customer,omitempty"`
	SaleDate        time.Time      `json:"sale_date"`
	Subtotal        float64        `gorm:"type:decimal(10,2);not null" json:"subtotal"`
	DiscountAmount  float64        `gorm:"type:decimal(15,2);default:0" json:"discount_amount"`
	TaxAmount       float64        `gorm:"type:decimal(15,2);default:0" json:"tax_amount"`
	TotalAmount     float64        `gorm:"type:decimal(15,2);not null" json:"total_amount"`
	FinalAmount     float64        `gorm:"type:decimal(15,2);not null" json:"final_amount"`
	PaymentMethodID *uint          `json:"payment_method_id"`
	PaymentMethod   *PaymentMethod `gorm:"foreignKey:PaymentMethodID" json:"payment_method,omitempty"`
	Status          string         `gorm:"size:20;not null" json:"status"` // 'pendente', 'pago', 'cancelado'
	Notes           string         `json:"notes"`
	CreatedByID     *uint          `gorm:"column:created_by" json:"created_by"`
	CreatedBy       *User          `gorm:"foreignKey:CreatedByID" json:"created_by_user,omitempty"`
	Items           []SaleItem     `gorm:"foreignKey:SaleID" json:"items,omitempty"`
}

// TableName especifica o nome da tabela
func (Sale) TableName() string {
	return "sales"
}
