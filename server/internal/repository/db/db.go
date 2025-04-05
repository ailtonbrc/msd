package repository

import (
	"clinica_server/internal/migrate"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConectarBanco(dsn string) *gorm.DB {
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("❌ Erro ao conectar ao banco de dados: %v", err)
    }

    log.Println("✅ Conectado ao banco com sucesso")

    // Executar automigrations de todas as entidades
    migrate.AutoMigrateUsuarios(db)
    migrate.AutoMigrateClinicas(db)
    migrate.AutoMigratePacientes(db)
    migrate.AutoMigrateRoles(db)

    return db
}