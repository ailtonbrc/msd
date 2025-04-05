package models

// CreateRoleRequest representa a estrutura de entrada para criação de role.
type CreateRoleRequest struct {
    Nome      string `json:"nome" binding:"required"`
    Descricao string `json:"descricao"`
}

// UpdateRoleRequest representa a estrutura de entrada para atualização de role.
type UpdateRoleRequest struct {
    Nome      string `json:"nome"`
    Descricao string `json:"descricao"`
}