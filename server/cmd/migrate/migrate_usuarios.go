package migrate

import "clinica_server/internal/models"
import "gorm.io/gorm"

func AutoMigrateUsuarios(db *gorm.DB) {
    db.AutoMigrate(&models.User{})
}