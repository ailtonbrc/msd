package migrate

import "clinica_server/internal/models"
import "gorm.io/gorm"

func AutoMigrateClinicas(db *gorm.DB) {
    db.AutoMigrate(&models.Clinica{})
}