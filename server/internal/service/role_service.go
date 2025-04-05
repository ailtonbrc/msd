package service

import (
    "clinica_server/internal/models"
    "gorm.io/gorm"
    "errors"
)

type RoleService struct {
    db *gorm.DB
}

func NewRoleService(db *gorm.DB) *RoleService {
    return &RoleService{db: db}
}

func (s *RoleService) ListarRoles() ([]models.Role, error) {
    var roles []models.Role
    err := s.db.Find(&roles).Error
    return roles, err
}

func (s *RoleService) BuscarRolePorID(id uint) (*models.Role, error) {
    var role models.Role
    if err := s.db.First(&role, id).Error; err != nil {
        return nil, err
    }
    return &role, nil
}

func (s *RoleService) CriarRole(req models.CreateRoleRequest) (*models.Role, error) {
    role := &models.Role{
        Nome:      req.Nome,
        Descricao: req.Descricao,
    }
    if err := s.db.Create(role).Error; err != nil {
        return nil, err
    }
    return role, nil
}

func (s *RoleService) AtualizarRole(id uint, req models.UpdateRoleRequest) (*models.Role, error) {
    role, err := s.BuscarRolePorID(id)
    if err != nil {
        return nil, err
    }
    role.Nome = req.Nome
    role.Descricao = req.Descricao
    if err := s.db.Save(role).Error; err != nil {
        return nil, err
    }
    return role, nil
}

func (s *RoleService) DeletarRole(id uint) error {
    return s.db.Delete(&models.Role{}, id).Error
}