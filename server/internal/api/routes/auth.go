package routes

import (
	"clinica_server/config"
	"clinica_server/internal/api/handlers"
	"clinica_server/internal/api/middlewares"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupAuthRoutes configura as rotas de autenticação
func SetupAuthRoutes(router *gin.RouterGroup, db *gorm.DB, cfg *config.Config) {
	log.Println("🔐 SetupAuthRoutes iniciado...") // ← Debug aqui
	
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
