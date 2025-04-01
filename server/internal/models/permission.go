package models

import (
	"gorm.io/gorm"
)

// Permission representa uma permissão no sistema
type Permission struct {
	gorm.Model

	Name        string `gorm:"size:100;not null;unique" json:"name"`
	Description string `json:"description"`
	Module      string `gorm:"size:50;not null" json:"module"`
	Roles       []Role `gorm:"many2many:role_permissions;" json:"-"`
}

// TableName especifica o nome da tabela
func (Permission) TableName() string {
	return "permissions"
}

// PermissionsByModule agrupa permissões por módulo
type PermissionsByModule struct {
	Module      string       `json:"module"`
	Permissions []Permission `json:"permissions"`
}
