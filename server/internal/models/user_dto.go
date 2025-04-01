package models

import (
	"simple-erp-service/internal/utils"
)

// UserDTO representa os dados de um usuário que são seguros para enviar ao frontend
type UserDTO struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email,omitempty"`
	RoleID   uint   `json:"role_id"`
	Role     string `json:"role,omitempty"` // Nome do perfil
	IsActive bool   `json:"is_active"`
}

// UserDetailDTO representa os dados detalhados de um usuário
type UserDetailDTO struct {
	ID        uint    `json:"id"`
	Username  string  `json:"username"`
	Name      string  `json:"name"`
	Email     string  `json:"email,omitempty"`
	RoleID    uint    `json:"role_id"`
	Role      RoleDTO `json:"role"`
	IsActive  bool    `json:"is_active"`
	LastLogin string  `json:"last_login,omitempty"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

// UserListDTO representa uma lista paginada de usuários
type UserListDTO struct {
	Users      []UserDTO      `json:"users"`
	Pagination *PaginationDTO `json:"pagination,omitempty"`
}

// PaginationDTO representa informações de paginação
type PaginationDTO struct {
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
	Sort       string `json:"sort"`
	Order      string `json:"order"`
	TotalRows  int64  `json:"total_rows"`
	TotalPages int    `json:"total_pages"`
}

// ToDTO converte um modelo User para UserDTO
func (u *User) ToDTO() UserDTO {
	dto := UserDTO{
		ID:       u.ID,
		Username: u.Username,
		Name:     u.Name,
		Email:    u.Email,
		RoleID:   u.RoleID,
		IsActive: u.IsActive,
	}

	// Adicionar o nome do perfil se estiver carregado
	if u.Role.ID != 0 {
		dto.Role = u.Role.Name
	}

	return dto
}

// ToDetailDTO converte um modelo User para UserDetailDTO
func (u *User) ToDetailDTO() UserDetailDTO {
	dto := UserDetailDTO{
		ID:        u.ID,
		Username:  u.Username,
		Name:      u.Name,
		Email:     u.Email,
		RoleID:    u.RoleID,
		IsActive:  u.IsActive,
		CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: u.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	// Adicionar o último login se existir
	if u.LastLogin != nil {
		dto.LastLogin = u.LastLogin.Format("2006-01-02 15:04:05")
	}

	// Adicionar o perfil se estiver carregado
	if u.Role.ID != 0 {
		dto.Role = u.Role.ToDTO()
	}

	return dto
}

// ToPaginationDTO converte um objeto de paginação para PaginationDTO
func ToPaginationDTO(pagination *utils.Pagination) *PaginationDTO {
	if pagination == nil {
		return nil
	}

	return &PaginationDTO{
		Page:       pagination.Page,
		Limit:      pagination.Limit,
		Sort:       pagination.Sort,
		Order:      pagination.Order,
		TotalRows:  pagination.TotalRows,
		TotalPages: pagination.TotalPages,
	}
}
