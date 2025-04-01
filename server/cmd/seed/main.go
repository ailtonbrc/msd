package main

import (
	"log"

	"simple-erp-service/config"
	"simple-erp-service/internal/repository/db"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// Carregar configurações
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar configurações: %v", err)
	}

	// Conectar ao banco de dados
	database, err := gorm.Open(postgres.Open(cfg.Database.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	// Executar apenas o seed
	if err := db.SeedDB(database); err != nil {
		log.Fatalf("Erro ao executar seed: %v", err)
	}

	log.Println("Seed concluído com sucesso!")
}
