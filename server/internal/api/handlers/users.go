package handlers

import (
	"net/http"
	"strconv"

	"simple-erp-service/internal/models"
	"simple-erp-service/internal/service"
	"simple-erp-service/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserHandler gerencia as requisições relacionadas a usuários
type UserHandler struct {
	userService *service.UserService
}

// NewUserHandler cria um novo handler de usuários
func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{
		userService: service.NewUserService(db),
	}
}

// GetUsers retorna uma lista paginada de usuários
// @Summary Listar usuários
// @Description Retorna uma lista paginada de usuários
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "Número da página" default(1)
// @Param limit query int false "Limite de itens por página" default(10)
// @Param sort query string false "Campo para ordenação" default(id)
// @Param order query string false "Direção da ordenação (asc/desc)" default(asc)
// @Success 200 {object} utils.Response "Usuários encontrados"
// @Failure 401 {object} utils.Response "Não autorizado"
// @Failure 500 {object} utils.Response "Erro ao buscar usuários"
// @Router /users [get]
func (h *UserHandler) GetUsers(c *gin.Context) {
	pagination := utils.GetPaginationParams(c)

	users, err := h.userService.GetUsers(&pagination)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Erro ao buscar usuários", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Usuários encontrados", users, nil)
}

// GetUser retorna um usuário específico
// @Summary Buscar usuário
// @Description Retorna um usuário específico pelo ID
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "ID do usuário"
// @Success 200 {object} utils.Response "Usuário encontrado"
// @Failure 400 {object} utils.Response "ID inválido"
// @Failure 401 {object} utils.Response "Não autorizado"
// @Failure 404 {object} utils.Response "Usuário não encontrado"
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "ID inválido", err.Error())
		return
	}

	user, err := h.userService.GetUserByID(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Usuário não encontrado", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Usuário encontrado", user, nil)
}

// CreateUser cria um novo usuário
// @Summary Criar usuário
// @Description Cria um novo usuário
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body models.CreateUserRequest true "Dados do usuário"
// @Success 201 {object} utils.Response "Usuário criado com sucesso"
// @Failure 400 {object} utils.Response "Dados inválidos"
// @Failure 401 {object} utils.Response "Não autorizado"
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, "Dados inválidos", err.Error())
		return
	}

	user, err := h.userService.CreateUser(req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Erro ao criar usuário", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Usuário criado com sucesso", user, nil)
}

// UpdateUser atualiza um usuário existente
// @Summary Atualizar usuário
// @Description Atualiza um usuário existente
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "ID do usuário"
// @Param request body models.UpdateUserRequest true "Dados do usuário"
// @Success 200 {object} utils.Response "Usuário atualizado com sucesso"
// @Failure 400 {object} utils.Response "Dados inválidos"
// @Failure 401 {object} utils.Response "Não autorizado"
// @Failure 404 {object} utils.Response "Usuário não encontrado"
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "ID inválido", err.Error())
		return
	}

	var req models.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, "Dados inválidos", err.Error())
		return
	}

	user, err := h.userService.UpdateUser(uint(id), req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Erro ao atualizar usuário", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Usuário atualizado com sucesso", user, nil)
}

// ChangePassword altera a senha de um usuário
// @Summary Alterar senha
// @Description Altera a senha de um usuário
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "ID do usuário"
// @Param request body models.ChangePasswordRequest true "Dados da senha"
// @Success 200 {object} utils.Response "Senha alterada com sucesso"
// @Failure 400 {object} utils.Response "Dados inválidos"
// @Failure 401 {object} utils.Response "Não autorizado"
// @Failure 404 {object} utils.Response "Usuário não encontrado"
// @Router /users/{id}/password [put]
func (h *UserHandler) ChangePassword(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "ID inválido", err.Error())
		return
	}

	var req models.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, "Dados inválidos", err.Error())
		return
	}

	// Verificar se o usuário é admin ou está alterando a própria senha
	userID, _ := c.Get("userID")
	role, _ := c.Get("role")
	isAdmin := role == "ADMIN"
	isSelf := userID.(uint) == uint(id)

	// Se não for admin e não for o próprio usuário, negar acesso
	if !isAdmin && !isSelf {
		utils.ErrorResponse(c, http.StatusForbidden, "Acesso negado", "Você não tem permissão para alterar a senha de outro usuário")
		return
	}

	// Se for admin alterando a senha de outro usuário, não precisa da senha atual
	if isAdmin && !isSelf {
		err = h.userService.ChangePassword(uint(id), "", req.NewPassword, true)
	} else {
		// Se for o próprio usuário ou admin alterando a própria senha, precisa da senha atual
		err = h.userService.ChangePassword(uint(id), req.CurrentPassword, req.NewPassword, false)
	}

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Erro ao alterar senha", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Senha alterada com sucesso", nil, nil)
}

// DeleteUser exclui um usuário
// @Summary Excluir usuário
// @Description Exclui um usuário
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "ID do usuário"
// @Success 200 {object} utils.Response "Usuário excluído com sucesso"
// @Failure 400 {object} utils.Response "ID inválido"
// @Failure 401 {object} utils.Response "Não autorizado"
// @Failure 404 {object} utils.Response "Usuário não encontrado"
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "ID inválido", err.Error())
		return
	}

	// Impedir que um usuário exclua a si mesmo
	userID, _ := c.Get("userID")
	if userID.(uint) == uint(id) {
		utils.ErrorResponse(c, http.StatusBadRequest, "Operação inválida", "Você não pode excluir seu próprio usuário")
		return
	}

	if err := h.userService.DeleteUser(uint(id)); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Erro ao excluir usuário", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Usuário excluído com sucesso", nil, nil)
}
