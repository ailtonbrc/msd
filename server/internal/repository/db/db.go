package db

import (
	"simple-erp-service/config"
	"simple-erp-service/migrations"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDB inicializa a conexão com o banco de dados
func InitDB(cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.Database.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	// Executar migrações
	if err := migrations.MigrateDB(db); err != nil {
		return nil, err
	}

	// Executar seed
	if err := SeedDB(db); err != nil {
		return nil, err
	}

	return db, nil
}
