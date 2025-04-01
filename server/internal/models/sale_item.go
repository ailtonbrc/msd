package models

import "gorm.io/gorm"

// SaleItem representa um item de venda
type SaleItem struct {
	gorm.Model

	SaleID          uint     `json:"sale_id"`
	Sale            *Sale    `gorm:"foreignKey:SaleID" json:"-"`
	ProductID       uint     `json:"product_id"`
	Product         *Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Quantity        int      `gorm:"not null" json:"quantity"`
	UnitPrice       float64  `gorm:"type:decimal(15,2);not null" json:"unit_price"`
	DiscountPercent float64  `gorm:"type:decimal(5,2);default:0" json:"discount_percent"`
	DiscountAmount  float64  `gorm:"type:decimal(15,2);default:0" json:"discount_amount"`
	TaxPercent      float64  `gorm:"type:decimal(5,2);default:0" json:"tax_percent"`
	TaxAmount       float64  `gorm:"type:decimal(15,2);default:0" json:"tax_amount"`
	TotalAmount     float64  `gorm:"type:decimal(15,2);not null" json:"total_amount"`
}

// TableName especifica o nome da tabela
func (SaleItem) TableName() string {
	return "sale_items"
}
