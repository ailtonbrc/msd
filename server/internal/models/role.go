package models

import "time"

// Role representa um papel/perfil de usuário no sistema.
type Role struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    Nome        string    `gorm:"size:100;not null;unique" json:"nome"`
    Descricao   string    `gorm:"size:255" json:"descricao"`
    CriadoEm    time.Time `gorm:"autoCreateTime" json:"criado_em"`
    AtualizadoEm time.Time `gorm:"autoUpdateTime" json:"atualizado_em"`
}

func (Role) TableName() string {
    return "roles"
}