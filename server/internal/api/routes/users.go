package routes

import (
	"simple-erp-service/config"
	"simple-erp-service/internal/api/handlers"
	"simple-erp-service/internal/api/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupUserRoutes configura as rotas de usuários
func SetupUserRoutes(router *gin.RouterGroup, db *gorm.DB) {
	userHandler := handlers.NewUserHandler(db)

	// Obter configuração para middleware de autenticação
	cfg, _ := config.Load()

	// Grupo de rotas de usuários (todas protegidas)
	users := router.Group("/users")
	users.Use(middlewares.AuthMiddleware(cfg))
	{
		users.GET("", middlewares.RequirePermission("users.view"), userHandler.GetUsers)
		users.GET("/:id", middlewares.RequirePermission("users.view"), userHandler.GetUser)
		users.POST("", middlewares.RequirePermission("users.create"), userHandler.CreateUser)
		users.PUT("/:id", middlewares.RequirePermission("users.edit"), userHandler.UpdateUser)
		users.DELETE("/:id", middlewares.RequirePermission("users.delete"), userHandler.DeleteUser)
		users.PUT("/:id/password", userHandler.ChangePassword) // Permissão verificada no handler
	}
}
