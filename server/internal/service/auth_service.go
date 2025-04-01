package service

import (
	"errors"
	"time"

	"simple-erp-service/config"
	"simple-erp-service/internal/models"
	"simple-erp-service/internal/utils"

	"gorm.io/gorm"
)

// AuthService gerencia a autenticação de usuários
type AuthService struct {
	db  *gorm.DB
	cfg *config.Config
}

// NewAuthService cria um novo serviço de autenticação
func NewAuthService(db *gorm.DB, cfg *config.Config) *AuthService {
	return &AuthService{
		db:  db,
		cfg: cfg,
	}
}

// LoginResponse representa a resposta do login
type LoginResponse struct {
	User         models.UserResponse `json:"user"`
	AccessToken  string              `json:"access_token"`
	RefreshToken string              `json:"refresh_token"`
	ExpiresIn    int                 `json:"expires_in"`
}

// Login autentica um usuário e retorna tokens JWT
func (s *AuthService) Login(username, password string) (*LoginResponse, error) {
	var user models.User

	// Buscar usuário pelo username
	result := s.db.Preload("Role.Permissions").Where("LOWER(username) = LOWER(?)", username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("usuário não encontrado")
		}
		return nil, result.Error
	}

	// Verificar se o usuário está ativo
	if !user.IsActive {
		return nil, errors.New("usuário inativo")
	}

	// Verificar senha
	if !utils.CheckPasswordHash(password, user.PasswordHash) {
		return nil, errors.New("senha incorreta")
	}

	// Extrair permissões
	var permissions []string
	if user.Role != nil {
		for _, perm := range user.Role.Permissions {
			permissions = append(permissions, perm.Name)
		}
	}

	// Gerar tokens
	accessToken, err := utils.GenerateAccessToken(user.ID, user.Username, user.RoleID, user.Role.Name, permissions, s.cfg)
	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID, user.Username, s.cfg)
	if err != nil {
		return nil, err
	}

	// Atualizar último login
	now := time.Now()
	user.LastLogin = &now
	s.db.Save(&user)

	return &LoginResponse{
		User:         user.ToResponse(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int(s.cfg.JWT.AccessTokenExp.Minutes()),
	}, nil
}

// RefreshToken renova o token de acesso usando um token de refresh
func (s *AuthService) RefreshToken(refreshToken string) (*LoginResponse, error) {
	// Validar token de refresh
	claims, err := utils.ValidateToken(refreshToken, s.cfg)
	if err != nil {
		return nil, err
	}

	// Buscar usuário Where("LOWER(username) = LOWER(?)", username).
	var user models.User
	result := s.db.Preload("Role.Permissions").Where("LOWER(username) = LOWER(?)", claims.Subject).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	// Verificar se o usuário está ativo
	if !user.IsActive {
		return nil, errors.New("usuário inativo")
	}

	// Extrair permissões
	var permissions []string
	if user.Role != nil {
		for _, perm := range user.Role.Permissions {
			permissions = append(permissions, perm.Name)
		}
	}

	// Gerar novo token de acesso
	newAccessToken, err := utils.GenerateAccessToken(user.ID, user.Username, user.RoleID, user.Role.Name, permissions, s.cfg)
	if err != nil {
		return nil, err
	}

	// Gerar novo token de refresh
	newRefreshToken, err := utils.GenerateRefreshToken(user.ID, user.Username, s.cfg)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		User:         user.ToResponse(),
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
		ExpiresIn:    int(s.cfg.JWT.AccessTokenExp.Minutes()),
	}, nil
}

// GetUserByID busca um usuário pelo ID
func (s *AuthService) GetUserByID(userID uint) (*models.User, error) {
	var user models.User
	result := s.db.Preload("Role.Permissions").First(&user, userID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
