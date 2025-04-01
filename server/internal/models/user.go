package models

import (
	"time"

	"gorm.io/gorm"
)

// User representa um usuário do sistema
type User struct {
	gorm.Model

	Username     string     `gorm:"size:50;not null;unique" json:"username"`
	PasswordHash string     `gorm:"size:255;not null;column:password_hash" json:"-"`
	Name         string     `gorm:"size:100;not null" json:"name"`
	Email        string     `gorm:"size:100;unique" json:"email"`
	Phone        string     `gorm:"size:20" json:"phone"`
	IsActive     bool       `gorm:"default:true" json:"is_active"`
	LastLogin    *time.Time `json:"last_login"`
	RoleID       uint       `json:"role_id"`
	Role         *Role      `gorm:"foreignKey:RoleID" json:"role,omitempty"`
}

// TableName especifica o nome da tabela
func (User) TableName() string {
	return "users"
}

// UserResponse é usado para retornar informações do usuário sem dados sensíveis
type UserResponse struct {
	ID        uint       `json:"id"`
	Username  string     `json:"username"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Phone     string     `json:"phone"`
	RoleID    uint       `json:"role_id"`
	Role      *Role      `json:"role,omitempty"`
	IsActive  bool       `json:"is_active"`
	LastLogin *time.Time `json:"last_login"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// ToResponse converte um User para UserResponse
func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:        u.ID,
		Username:  u.Username,
		Name:      u.Name,
		Email:     u.Email,
		Phone:     u.Phone,
		RoleID:    u.RoleID,
		Role:      u.Role,
		IsActive:  u.IsActive,
		LastLogin: u.LastLogin,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

// CreateUserRequest representa os dados para criar um novo usuário
type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=6"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone"`
	RoleID   uint   `json:"role_id" binding:"required"`
}

// UpdateUserRequest representa os dados para atualizar um usuário
type UpdateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email" binding:"omitempty,email"`
	Phone    string `json:"phone"`
	RoleID   uint   `json:"role_id"`
	IsActive *bool  `json:"is_active"`
}

// ChangePasswordRequest representa os dados para alterar a senha
type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required,min=6"`
}
