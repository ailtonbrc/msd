package models

import (
	"time"

	"gorm.io/gorm"
)

// Purchase representa uma compra
type Purchase struct {
	gorm.Model

	SupplierID   *uint          `json:"supplier_id"`
	Supplier     *Supplier      `gorm:"foreignKey:SupplierID" json:"supplier,omitempty"`
	PurchaseDate time.Time      `json:"purchase_date"`
	TotalAmount  float64        `gorm:"type:decimal(15,2);not null" json:"total_amount"`
	Status       string         `gorm:"size:20;not null" json:"status"` // 'pendente', 'recebido', 'cancelado'
	Notes        string         `json:"notes"`
	CreatedByID  *uint          `gorm:"column:created_by" json:"created_by"`
	CreatedBy    *User          `gorm:"foreignKey:CreatedByID" json:"created_by_user,omitempty"`
	Items        []PurchaseItem `gorm:"foreignKey:PurchaseID" json:"items,omitempty"`
}

// TableName especifica o nome da tabela
func (Purchase) TableName() string {
	return "purchases"
}
