package main

import (
	"log"

	"simple-erp-service/config"
	"simple-erp-service/internal/repository/db"
)

func main() {
	// Carregar configurações
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar configurações: %v", err)
	}

	// Conectar ao banco de dados
	_, err = db.InitDB(cfg)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	log.Println("Migrações e seed concluídos com sucesso!")
}
