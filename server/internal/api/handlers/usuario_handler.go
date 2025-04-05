package handlers

import (
    "net/http"
    "strconv"

    "clinica_server/internal/models"
    "clinica_server/internal/service"
    "clinica_server/internal/utils"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type UsuarioHandler struct {
    usuarioService *service.UsuarioService
}

func NewUsuarioHandler(db *gorm.DB) *UsuarioHandler {
    return &UsuarioHandler{
        usuarioService: service.NewUsuarioService(db),
    }
}

func (h *UsuarioHandler) GetUsuarios(c *gin.Context) {
    pagination := utils.GetPaginationParams(c)
    items, err := h.usuarioService.GetUsuarios(&pagination)
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, "Erro ao buscar usuarios", err.Error())
        return
    }
    utils.SuccessResponse(c, http.StatusOK, "Usuarios encontrados", items, nil)
}

func (h *UsuarioHandler) GetUsuario(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, "ID inválido", err.Error())
        return
    }
    item, err := h.usuarioService.GetUsuarioByID(uint(id))
    if err != nil {
        utils.ErrorResponse(c, http.StatusNotFound, "Usuario não encontrado", err.Error())
        return
    }
    utils.SuccessResponse(c, http.StatusOK, "Usuario encontrado", item, nil)
}

func (h *UsuarioHandler) CreateUsuario(c *gin.Context) {
    var req models.CreateUsuarioRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        utils.ValidationErrorResponse(c, "Dados inválidos", err.Error())
        return
    }
    item, err := h.usuarioService.CreateUsuario(req)
    if err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, "Erro ao criar usuario", err.Error())
        return
    }
    utils.SuccessResponse(c, http.StatusCreated, "Usuario criado com sucesso", item, nil)
}

func (h *UsuarioHandler) UpdateUsuario(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, "ID inválido", err.Error())
        return
    }
    var req models.UpdateUsuarioRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        utils.ValidationErrorResponse(c, "Dados inválidos", err.Error())
        return
    }
    item, err := h.usuarioService.UpdateUsuario(uint(id), req)
    if err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, "Erro ao atualizar usuario", err.Error())
        return
    }
    utils.SuccessResponse(c, http.StatusOK, "Usuario atualizado com sucesso", item, nil)
}

func (h *UsuarioHandler) DeleteUsuario(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, "ID inválido", err.Error())
        return
    }
    if err := h.usuarioService.DeleteUsuario(uint(id)); err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, "Erro ao excluir usuario", err.Error())
        return
    }
    utils.SuccessResponse(c, http.StatusOK, "Usuario excluído com sucesso", nil, nil)
}