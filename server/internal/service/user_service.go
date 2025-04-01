package service

import (
	"errors"

	"simple-erp-service/internal/models"
	"simple-erp-service/internal/utils"

	"gorm.io/gorm"
)

// UserService gerencia operações relacionadas a usuários
type UserService struct {
	db *gorm.DB
}

// NewUserService cria um novo serviço de usuários
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
}

// GetUsers retorna uma lista paginada de usuários
func (s *UserService) GetUsers(pagination *utils.Pagination) (*models.UserListDTO, error) {
	var users []models.User

	query := s.db.Model(&models.User{}).Preload("Role")
	query, err := utils.Paginate(&models.User{}, pagination, query)
	if err != nil {
		return nil, err
	}

	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}

	// Converter para DTOs
	userDTOs := make([]models.UserDTO, 0, len(users))
	for _, user := range users {
		userDTOs = append(userDTOs, user.ToDTO())
	}

	return &models.UserListDTO{
		Users:      userDTOs,
		Pagination: models.ToPaginationDTO(pagination),
	}, nil
}

// GetUserByID busca um usuário pelo ID
func (s *UserService) GetUserByID(id uint) (*models.UserDetailDTO, error) {
	var user models.User
	if err := s.db.Preload("Role").First(&user, id).Error; err != nil {
		return nil, err
	}

	// Converter para DTO
	userDetailDTO := user.ToDetailDTO()
	return &userDetailDTO, nil
}

// CreateUser cria um novo usuário
func (s *UserService) CreateUser(req models.CreateUserRequest) (*models.UserDTO, error) {
	// Verificar se o username já existe
	var count int64
	s.db.Model(&models.User{}).Where("username = ?", req.Username).Count(&count)
	if count > 0 {
		return nil, errors.New("nome de usuário já está em uso")
	}

	// Verificar se o email já existe (se fornecido)
	if req.Email != "" {
		s.db.Model(&models.User{}).Where("email = ?", req.Email).Count(&count)
		if count > 0 {
			return nil, errors.New("email já está em uso")
		}
	}

	// Verificar se o perfil existe
	var role models.Role
	if err := s.db.First(&role, req.RoleID).Error; err != nil {
		return nil, errors.New("perfil não encontrado")
	}

	// Hash da senha
	passwordHash, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// Criar usuário
	user := models.User{
		Username:     req.Username,
		PasswordHash: passwordHash,
		Name:         req.Name,
		Email:        req.Email,
		RoleID:       req.RoleID,
		IsActive:     true, // Por padrão, usuários são criados ativos
	}

	if err := s.db.Create(&user).Error; err != nil {
		return nil, err
	}

	// Carregar o perfil para o DTO
	s.db.Preload("Role").First(&user, user.ID)

	// Converter para DTO
	userDTO := user.ToDTO()
	return &userDTO, nil
}

// UpdateUser atualiza um usuário existente
func (s *UserService) UpdateUser(id uint, req models.UpdateUserRequest) (*models.UserDTO, error) {
	// Buscar usuário
	var user models.User
	if err := s.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	// Verificar se o email já está em uso por outro usuário (se fornecido)
	if req.Email != "" && req.Email != user.Email {
		var count int64
		s.db.Model(&models.User{}).Where("email = ? AND id != ?", req.Email, id).Count(&count)
		if count > 0 {
			return nil, errors.New("email já está em uso")
		}
	}

	// Verificar se o perfil existe (se fornecido)
	if req.RoleID != 0 && req.RoleID != user.RoleID {
		var role models.Role
		if err := s.db.First(&role, req.RoleID).Error; err != nil {
			return nil, errors.New("perfil não encontrado")
		}
	}

	// Atualizar campos
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.RoleID != 0 {
		user.RoleID = req.RoleID
	}
	if req.IsActive != nil {
		user.IsActive = *req.IsActive
	}

	// Salvar alterações
	if err := s.db.Save(&user).Error; err != nil {
		return nil, err
	}

	// Carregar o perfil para o DTO
	s.db.Preload("Role").First(&user, user.ID)

	// Converter para DTO
	userDTO := user.ToDTO()
	return &userDTO, nil
}

// ChangePassword altera a senha de um usuário
func (s *UserService) ChangePassword(id uint, currentPassword, newPassword string, isAdmin bool) error {
	// Buscar usuário
	var user models.User
	if err := s.db.First(&user, id).Error; err != nil {
		return err
	}

	// Se não for admin, verificar a senha atual
	if !isAdmin {
		if !utils.CheckPasswordHash(currentPassword, user.PasswordHash) {
			return errors.New("senha atual incorreta")
		}
	}

	// Hash da nova senha
	passwordHash, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}

	// Atualizar senha
	user.PasswordHash = passwordHash
	return s.db.Save(&user).Error
}

// DeleteUser exclui um usuário (soft delete)
func (s *UserService) DeleteUser(id uint) error {
	return s.db.Delete(&models.User{}, id).Error
}
