package models

// Permission define uma permissão específica no sistema
type Permission struct {
    ID     uint   `gorm:"primaryKey" json:"id"`
    Codigo string `gorm:"unique;not null" json:"codigo"`
    Nome   string `json:"nome"`
}