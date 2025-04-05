package models

import (
    "gorm.io/gorm"
    "time"
)

type Usuario struct {
    ID                    uint           `gorm:"primaryKey" json:"id"`
    Nome                  string         `json:"nome"`
    Email                 string         `gorm:"unique" json:"email"`
    Senha                 string         `json:"senha"`
    Perfil                string         `json:"perfil"`
    ClinicaID             *uint          `json:"clinica_id"`
    SupervisorID          *uint          `json:"supervisor_id"`
    Ativo                 bool           `gorm:"default:true" json:"ativo"`
    DataInicioInatividade *time.Time     `json:"data_inicio_inatividade"`
    DataFimInatividade    *time.Time     `json:"data_fim_inatividade"`
    MotivoInatividade     *string        `json:"motivo_inatividade"`
    CriadoEm              time.Time      `gorm:"autoCreateTime" json:"criado_em"`
    AtualizadoEm          time.Time      `gorm:"autoUpdateTime" json:"atualizado_em"`
    DeletedAt             gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Usuario) TableName() string {
    return "usuarios"
}