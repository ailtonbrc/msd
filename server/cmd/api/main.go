package main

import (
	"log"

	"simple-erp-service/config"
	"simple-erp-service/internal/api/server"
	"simple-erp-service/internal/repository/db"
)

func main() {
	// Carregar configurações
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar configurações: %v", err)
	}

	// Inicializar banco de dados
	db, err := db.InitDB(cfg)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	// Inicializar e executar o servidor
	s := server.NewServer(cfg, db)
	if err := s.Run(); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
