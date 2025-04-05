package routes

import (
    "clinica_server/internal/api/handlers"
    "clinica_server/internal/api/middlewares"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func SetupRoleRoutes(rg *gin.RouterGroup, db *gorm.DB) {
    roleHandler := handlers.NewRoleHandler(db)
    protected := rg.Group("/roles")
    protected.Use(middlewares.AuthMiddleware())
    {
        protected.GET("", roleHandler.List)
        protected.GET("/:id", roleHandler.Get)
        protected.POST("", roleHandler.Create)
        protected.PUT("/:id", roleHandler.Update)
        protected.DELETE("/:id", roleHandler.Delete)
    }
}