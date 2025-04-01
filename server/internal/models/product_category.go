package models

import "gorm.io/gorm"

// ProductCategory representa uma categoria de produto
type ProductCategory struct {
	gorm.Model

	Name        string            `gorm:"size:100;not null" json:"name"`
	Description string            `json:"description"`
	ParentID    *uint             `json:"parent_id"`
	Parent      *ProductCategory  `gorm:"foreignKey:ParentID" json:"parent,omitempty"`
	Children    []ProductCategory `gorm:"foreignKey:ParentID" json:"children,omitempty"`
	Products    []Product         `gorm:"foreignKey:CategoryID" json:"-"`
}

// TableName especifica o nome da tabela
func (ProductCategory) TableName() string {
	return "product_categories"
}
