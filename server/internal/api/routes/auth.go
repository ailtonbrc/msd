package routes

import (
	"simple-erp-service/config"
	"simple-erp-service/internal/api/handlers"
	"simple-erp-service/internal/api/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupAuthRoutes configura as rotas de autenticação
func SetupAuthRoutes(router *gin.RouterGroup, db *gorm.DB, cfg *config.Config) {
	authHandler := handlers.NewAuthHandler(db, cfg)

	auth := router.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
		auth.POST("/refresh-token", authHandler.RefreshToken)

		// Rotas protegidas
		protected := auth.Group("")
		protected.Use(middlewares.AuthMiddleware(cfg))
		{
			protected.POST("/logout", authHandler.Logout)
			protected.GET("/me", authHandler.GetMe)
		}
	}
}
