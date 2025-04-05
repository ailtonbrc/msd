package repository

import (
    "clinica_server/internal/models"
    "gorm.io/gorm"
)

type RoleRepository struct {
    DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
    return &RoleRepository{DB: db}
}

func (r *RoleRepository) FindAll() ([]models.Role, error) {
    var roles []models.Role
    err := r.DB.Find(&roles).Error
    return roles, err
}