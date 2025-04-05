package migrate

import (
    "clinica_server/internal/models"
    "gorm.io/gorm"
    "log"
)

func AutoMigrateRoles(db *gorm.DB) {
    log.Println("📦 Migrando tabelas de roles e permissões...")
    err := db.AutoMigrate(&models.Role{}, &models.Permission{})
    if err != nil {
        log.Fatalf("❌ Falha ao migrar Role/Permission: %v", err)
    }
    log.Println("✅ Migração de Role e Permission concluída.")
}