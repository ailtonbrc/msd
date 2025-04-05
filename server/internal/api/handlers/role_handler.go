package handlers

import (
    "clinica_server/internal/models"
    "clinica_server/internal/service"
    "clinica_server/internal/utils"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type RoleHandler struct {
    service *service.RoleService
}

func NewRoleHandler(db *gorm.DB) *RoleHandler {
    return &RoleHandler{service: service.NewRoleService(db)}
}

func (h *RoleHandler) RegisterRoutes(rg *gin.RouterGroup) {
    roles := rg.Group("/roles")
    roles.GET("", h.List)
    roles.GET("/:id", h.Get)
    roles.POST("", h.Create)
    roles.PUT("/:id", h.Update)
    roles.DELETE("/:id", h.Delete)
}

func (h *RoleHandler) List(c *gin.Context) {
    items, err := h.service.ListarRoles()
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, "Erro ao listar roles", err.Error())
        return
    }
    utils.SuccessResponse(c, http.StatusOK, "Roles encontrados", items, nil)
}

func (h *RoleHandler) Get(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    item, err := h.service.BuscarRolePorID(uint(id))
    if err != nil || item == nil {
        utils.ErrorResponse(c, http.StatusNotFound, "Role não encontrada", err)
        return
    }
    utils.SuccessResponse(c, http.StatusOK, "Role encontrada", item, nil)
}

func (h *RoleHandler) Create(c *gin.Context) {
    var req models.CreateRoleRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        utils.ValidationErrorResponse(c, "Dados inválidos", err.Error())
        return
    }
    item, err := h.service.CriarRole(req)
    if err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, "Erro ao criar role", err.Error())
        return
    }
    utils.SuccessResponse(c, http.StatusCreated, "Role criada", item, nil)
}

func (h *RoleHandler) Update(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var req models.UpdateRoleRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        utils.ValidationErrorResponse(c, "Dados inválidos", err.Error())
        return
    }
    item, err := h.service.AtualizarRole(uint(id), req)
    if err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, "Erro ao atualizar role", err.Error())
        return
    }
    utils.SuccessResponse(c, http.StatusOK, "Role atualizada", item, nil)
}

func (h *RoleHandler) Delete(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    if err := h.service.DeletarRole(uint(id)); err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, "Erro ao excluir role", err.Error())
        return
    }
    utils.SuccessResponse(c, http.StatusOK, "Role excluída", nil, nil)
}