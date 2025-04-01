package models

import "gorm.io/gorm"

// PurchaseItem representa um item de compra
type PurchaseItem struct {
	gorm.Model

	PurchaseID  uint      `json:"purchase_id"`
	Purchase    *Purchase `gorm:"foreignKey:PurchaseID" json:"-"`
	ProductID   uint      `json:"product_id"`
	Product     *Product  `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Quantity    int       `gorm:"not null" json:"quantity"`
	UnitPrice   float64   `gorm:"type:decimal(15,2);not null" json:"unit_price"`
	TotalAmount float64   `gorm:"type:decimal(15,2);not null" json:"total_amount"`
}

// TableName especifica o nome da tabela
func (PurchaseItem) TableName() string {
	return "purchase_items"
}
