package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/msd/server/handlers"
)

func UsuarioRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/usuarios", handlers.CriarUsuario)
}
