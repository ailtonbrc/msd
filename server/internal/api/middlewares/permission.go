package middlewares

import (
	"net/http"

	"simple-erp-service/internal/utils"

	"github.com/gin-gonic/gin"
)

// RequirePermission verifica se o usuário tem a permissão necessária
func RequirePermission(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Verificar se o usuário está autenticado
		permissions, exists := c.Get("permissions")
		if !exists {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Não autorizado", "Usuário não autenticado")
			c.Abort()
			return
		}

		// Verificar se o usuário tem a permissão necessária
		userPermissions, ok := permissions.([]string)
		if !ok {
			utils.ErrorResponse(c, http.StatusInternalServerError, "Erro interno", "Erro ao verificar permissões")
			c.Abort()
			return
		}

		// Verificar se o usuário é admin (tem todas as permissões)
		role, exists := c.Get("role")
		if exists && role.(string) == "ADMIN" {
			c.Next()
			return
		}

		// Verificar se o usuário tem a permissão específica
		hasPermission := false
		for _, p := range userPermissions {
			if p == permission {
				hasPermission = true
				break
			}
		}

		if !hasPermission {
			utils.ErrorResponse(c, http.StatusForbidden, "Acesso negado", "Permissão insuficiente")
			c.Abort()
			return
		}

		c.Next()
	}
}
