package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/ailtonbrc/msd/server/database"
	"github.com/ailtonbrc/msd/server/middleware"
	"github.com/ailtonbrc/msd/server/routes"
)

func main() {
	// Carregar variáveis de ambiente
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o .env")
	}

	// Conectar ao banco de dados
	database.Conectar()
	defer database.DB.Close()

	// Criar app Fiber
	app := fiber.New()

	// Rotas básicas
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API Go está rodando! 🚀")
	})

	app.Get("/api/teste", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"mensagem": "Rota de teste funcionando com sucesso!",
			"status":   "ok",
			"hora":     time.Now().Format(time.RFC3339),
		})
	})

	

	// Agrupar rotas da API
	api := app.Group("/api")

	// Registrar rotas 
	log.Println("Registrando SetupAuthRoutes")
	routes.SetupAuthRoutes(api)
	log.Println("Registrando SetupUsuarioRoutes")
	routes.SetupUsuarioRoutes(api)
	routes.SetupClinicaRoutes(api)

	// ✅ Aqui adicionamos a rota protegida com middleware JWT
	api.Get("/rota-protegida", middleware.AuthRequired, func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"mensagem": "Você está autenticado!"})
	})

	// Porta do servidor
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	log.Println("✅ Servidor iniciado na porta " + port)
	log.Fatal(app.Listen(":" + port))
}
