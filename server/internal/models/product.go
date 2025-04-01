package models

import "gorm.io/gorm"

// Product representa um produto
type Product struct {
	gorm.Model

	SKU          string           `gorm:"size:50;unique" json:"sku"`
	Barcode      string           `gorm:"size:50;unique" json:"barcode"`
	Name         string           `gorm:"size:255;not null" json:"name"`
	Description  string           `json:"description"`
	CategoryID   *uint            `json:"category_id"`
	Category     *ProductCategory `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	UnitID       *uint            `json:"unit_id"`
	Unit         *MeasurementUnit `gorm:"foreignKey:UnitID" json:"unit,omitempty"`
	CostPrice    float64          `gorm:"type:decimal(15,2);not null" json:"cost_price"`
	SellingPrice float64          `gorm:"type:decimal(15,2);not null" json:"selling_price"`
	MinStock     int              `gorm:"default:0" json:"min_stock"`
	MaxStock     *int             `json:"max_stock"`
	CurrentStock int              `gorm:"default:0" json:"current_stock"`
	IsActive     bool             `gorm:"default:true" json:"is_active"`
	CreatedByID  *uint            `gorm:"column:created_by" json:"created_by"`
	CreatedBy    *User            `gorm:"foreignKey:CreatedByID" json:"created_by_user,omitempty"`
}

// TableName especifica o nome da tabela
func (Product) TableName() string {
	return "products"
}
