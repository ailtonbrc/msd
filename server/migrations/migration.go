package migrations

import (
	"log"
	"simple-erp-service/internal/models"

	"gorm.io/gorm"
)

// MigrateDB executa as migrações do banco de dados
func MigrateDB(db *gorm.DB) error {
	log.Println("Iniciando migrações do banco de dados...")

	// Lista de todos os modelos para migração
	models := []interface{}{
		&models.Role{},
		&models.Permission{},
		&models.User{},
		&models.ProductCategory{},
		&models.MeasurementUnit{},
		&models.Product{},
		&models.InventoryMovement{},
		&models.Customer{},
		&models.Supplier{},
		&models.PaymentMethod{},
		&models.Sale{},
		&models.SaleItem{},
		&models.Purchase{},
		&models.PurchaseItem{},
		&models.FinancialTransaction{},
		&models.SystemLog{},
	}

	// Executar migrações
	err := db.AutoMigrate(models...)
	if err != nil {
		log.Printf("Erro ao executar migrações: %v", err)
		return err
	}

	log.Println("Migrações concluídas com sucesso!")
	return nil
}
