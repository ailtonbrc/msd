package migrate

import "clinica_server/internal/models"
import "gorm.io/gorm"

func AutoMigratePacientes(db *gorm.DB) {
    db.AutoMigrate(&models.Paciente{})
}