package models

// RoleDTO representa os dados de um perfil que são seguros para enviar ao frontend
type RoleDTO struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	// Não incluímos campos sensíveis ou desnecessários aqui
}

// RoleDetailDTO representa os dados detalhados de um perfil, incluindo suas permissões
type RoleDetailDTO struct {
	ID          uint            `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Permissions []PermissionDTO `json:"permissions"`
	CreatedAt   string          `json:"created_at"`
	UpdatedAt   string          `json:"updated_at"`
}

// PermissionDTO representa os dados de uma permissão que são seguros para enviar ao frontend
type PermissionDTO struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Module      string `json:"module"`
}

// PermissionsByModuleDTO representa permissões agrupadas por módulo
type PermissionsByModuleDTO struct {
	Module      string          `json:"module"`
	Permissions []PermissionDTO `json:"permissions"`
}

// ToDTO converte um modelo Role para RoleDTO
func (r *Role) ToDTO() RoleDTO {
	return RoleDTO{
		ID:          r.ID,
		Name:        r.Name,
		Description: r.Description,
	}
}

// ToDetailDTO converte um modelo Role para RoleDetailDTO
func (r *Role) ToDetailDTO() RoleDetailDTO {
	permissionDTOs := make([]PermissionDTO, 0, len(r.Permissions))
	for _, perm := range r.Permissions {
		permissionDTOs = append(permissionDTOs, perm.ToDTO())
	}

	return RoleDetailDTO{
		ID:          r.ID,
		Name:        r.Name,
		Description: r.Description,
		Permissions: permissionDTOs,
		CreatedAt:   r.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   r.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

// ToDTO converte um modelo Permission para PermissionDTO
func (p *Permission) ToDTO() PermissionDTO {
	return PermissionDTO{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Module:      p.Module,
	}
}

// ToDTO converte um modelo PermissionsByModule para PermissionsByModuleDTO
func (p *PermissionsByModule) ToDTO() PermissionsByModuleDTO {
	permissionDTOs := make([]PermissionDTO, 0, len(p.Permissions))
	for _, perm := range p.Permissions {
		permissionDTOs = append(permissionDTOs, perm.ToDTO())
	}

	return PermissionsByModuleDTO{
		Module:      p.Module,
		Permissions: permissionDTOs,
	}
}