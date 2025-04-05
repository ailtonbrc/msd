package main

import (
	"clinica_server/config"
	"clinica_server/internal/api/routes"
	"clinica_server/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {
    cfg, _ := config.Load()

    //✅ Conecta ao PostgreSQL
    db := repository.ConectarBanco(cfg.Database.DSN()) 

    r := gin.Default()

    api := r.Group("/api")
    routes.SetupAuthRoutes(api, db, &cfg)
    routes.SetupUsuarioRoutes(api, db)
    routes.SetupClinicaRoutes(api, db)
    routes.SetupPacienteRoutes(api, db)
    routes.SetupRoleRoutes(api, db)

    r.Run(":" + cfg.Server.Port)
}