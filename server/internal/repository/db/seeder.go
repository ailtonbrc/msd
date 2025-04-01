package db

import (
	"log"

	"simple-erp-service/internal/models"
	"simple-erp-service/internal/utils"

	"gorm.io/gorm"
)

// SeedDB popula o banco de dados com dados iniciais
func SeedDB(db *gorm.DB) error {
	log.Println("Iniciando seed do banco de dados...")

	// Verificar se já existem dados
	var count int64
	db.Model(&models.Role{}).Count(&count)
	if count > 0 {
		log.Println("Banco de dados já possui dados. Pulando seed.")
		return nil
	}

	// Iniciar uma transação
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Seed de Roles (Perfis)
	roles := []models.Role{
		{Name: "ADMIN", Description: "Administrador do sistema com acesso completo"},
		{Name: "Gerente", Description: "Acesso gerencial a múltiplos módulos"},
		{Name: "Vendas", Description: "Acesso ao módulo de vendas"},
		{Name: "Estoque", Description: "Acesso ao módulo de estoque"},
		{Name: "Financeiro", Description: "Acesso ao módulo financeiro"},
	}

	if err := tx.Create(&roles).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Seed de Permissions (Permissões)
	permissions := []models.Permission{
		// Módulo de Usuários
		{Name: "users.view", Description: "Visualizar usuários", Module: "users"},
		{Name: "users.create", Description: "Criar usuários", Module: "users"},
		{Name: "users.edit", Description: "Editar usuários", Module: "users"},
		{Name: "users.delete", Description: "Excluir usuários", Module: "users"},

		// Módulo de Vendas
		{Name: "sales.view", Description: "Visualizar vendas", Module: "sales"},
		{Name: "sales.create", Description: "Criar vendas", Module: "sales"},
		{Name: "sales.edit", Description: "Editar vendas", Module: "sales"},
		{Name: "sales.delete", Description: "Excluir vendas", Module: "sales"},
		{Name: "sales.reports", Description: "Gerar relatórios de vendas", Module: "sales"},

		// Módulo de Estoque
		{Name: "inventory.view", Description: "Visualizar estoque", Module: "inventory"},
		{Name: "inventory.create", Description: "Adicionar itens ao estoque", Module: "inventory"},
		{Name: "inventory.edit", Description: "Editar itens do estoque", Module: "inventory"},
		{Name: "inventory.delete", Description: "Remover itens do estoque", Module: "inventory"},
		{Name: "inventory.reports", Description: "Gerar relatórios de estoque", Module: "inventory"},

		// Módulo Financeiro
		{Name: "finance.view", Description: "Visualizar finanças", Module: "finance"},
		{Name: "finance.create", Description: "Criar transações financeiras", Module: "finance"},
		{Name: "finance.edit", Description: "Editar transações financeiras", Module: "finance"},
		{Name: "finance.delete", Description: "Excluir transações financeiras", Module: "finance"},
		{Name: "finance.reports", Description: "Gerar relatórios financeiros", Module: "finance"},
	}

	if err := tx.Create(&permissions).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Atribuir todas as permissões ao ADMIN
	var adminRole models.Role
	if err := tx.Where("name = ?", "ADMIN").First(&adminRole).Error; err != nil {
		tx.Rollback()
		return err
	}

	var allPermissions []models.Permission
	if err := tx.Find(&allPermissions).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&adminRole).Association("Permissions").Append(&allPermissions); err != nil {
		tx.Rollback()
		return err
	}

	// Atribuir permissões ao perfil de Vendas
	var salesRole models.Role
	if err := tx.Where("name = ?", "Vendas").First(&salesRole).Error; err != nil {
		tx.Rollback()
		return err
	}

	var salesPermissions []models.Permission
	if err := tx.Where("module = ?", "sales").Find(&salesPermissions).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&salesRole).Association("Permissions").Append(&salesPermissions); err != nil {
		tx.Rollback()
		return err
	}

	// Atribuir permissões ao perfil de Estoque
	var inventoryRole models.Role
	if err := tx.Where("name = ?", "Estoque").First(&inventoryRole).Error; err != nil {
		tx.Rollback()
		return err
	}

	var inventoryPermissions []models.Permission
	if err := tx.Where("module = ?", "inventory").Find(&inventoryPermissions).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&inventoryRole).Association("Permissions").Append(&inventoryPermissions); err != nil {
		tx.Rollback()
		return err
	}

	// Atribuir permissões ao perfil de Financeiro
	var financeRole models.Role
	if err := tx.Where("name = ?", "Financeiro").First(&financeRole).Error; err != nil {
		tx.Rollback()
		return err
	}

	var financePermissions []models.Permission
	if err := tx.Where("module = ?", "finance").Find(&financePermissions).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&financeRole).Association("Permissions").Append(&financePermissions); err != nil {
		tx.Rollback()
		return err
	}

	// Atribuir permissões de visualização de todos os módulos ao Gerente
	var managerRole models.Role
	if err := tx.Where("name = ?", "Gerente").First(&managerRole).Error; err != nil {
		tx.Rollback()
		return err
	}

	var viewPermissions []models.Permission
	if err := tx.Where("name LIKE ?", "%.view").Or("name LIKE ?", "%.reports").Find(&viewPermissions).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&managerRole).Association("Permissions").Append(&viewPermissions); err != nil {
		tx.Rollback()
		return err
	}

	// Criar usuário admin
	passwordHash, err := utils.HashPassword("987321")
	if err != nil {
		tx.Rollback()
		return err
	}

	adminUser := models.User{
		Username:     "admin",
		PasswordHash: passwordHash,
		Name:         "Administrador",
		Email:        "admin@sistema.com",
		RoleID:       adminRole.ID,
		IsActive:     true,
	}

	if err := tx.Create(&adminUser).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Seed de unidades de medida
	units := []models.MeasurementUnit{
		{Name: "Unidade", Abbreviation: "un"},
		{Name: "Caixa", Abbreviation: "cx"},
		{Name: "Pacote", Abbreviation: "pct"},
		{Name: "Quilograma", Abbreviation: "kg"},
		{Name: "Litro", Abbreviation: "l"},
		{Name: "Metro", Abbreviation: "m"},
	}

	if err := tx.Create(&units).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Seed de métodos de pagamento
	paymentMethods := []models.PaymentMethod{
		{Name: "Dinheiro", Description: "Pagamento em espécie", IsActive: true},
		{Name: "Cartão de Crédito", Description: "Pagamento com cartão de crédito", IsActive: true},
		{Name: "Cartão de Débito", Description: "Pagamento com cartão de débito", IsActive: true},
		{Name: "Transferência Bancária", Description: "Pagamento via transferência bancária", IsActive: true},
		{Name: "Pix", Description: "Pagamento via Pix", IsActive: true},
		{Name: "Boleto", Description: "Pagamento via boleto bancário", IsActive: true},
	}

	if err := tx.Create(&paymentMethods).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Seed de categorias de produtos
	categories := []models.ProductCategory{
		{Name: "Geral", Description: "Categoria geral de produtos"},
		{Name: "Eletrônicos", Description: "Produtos eletrônicos"},
		{Name: "Alimentos", Description: "Produtos alimentícios"},
		{Name: "Vestuário", Description: "Roupas e acessórios"},
		{Name: "Papelaria", Description: "Material de escritório"},
	}

	if err := tx.Create(&categories).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit da transação
	if err := tx.Commit().Error; err != nil {
		return err
	}

	log.Println("Seed do banco de dados concluído com sucesso!")
	return nil
}
