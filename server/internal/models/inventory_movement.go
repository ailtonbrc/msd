package models

import "gorm.io/gorm"

// InventoryMovement representa uma movimentação de estoque
type InventoryMovement struct {
	gorm.Model

	ProductID     uint     `json:"product_id"`
	Product       *Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Quantity      int      `gorm:"not null" json:"quantity"`
	PreviousStock int      `gorm:"not null" json:"previous_stock"`
	NewStock      int      `gorm:"not null" json:"new_stock"`
	MovementType  string   `gorm:"size:20;not null" json:"movement_type"` // 'entrada', 'saida', 'ajuste'
	ReferenceID   *uint    `json:"reference_id"`                          // ID da venda, compra ou ajuste
	ReferenceType string   `gorm:"size:20" json:"reference_type"`         // 'venda', 'compra', 'ajuste'
	Notes         string   `json:"notes"`
	CreatedByID   *uint    `gorm:"column:created_by" json:"created_by"`
	CreatedBy     *User    `gorm:"foreignKey:CreatedByID" json:"created_by_user,omitempty"`
}

// TableName especifica o nome da tabela
func (InventoryMovement) TableName() string {
	return "inventory_movements"
}
