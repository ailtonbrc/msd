package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"clinica_server/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

// InitDB inicializa a conexão com o banco de dados via GORM e aplica o schema inicial se necessário.
func InitDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := cfg.Database.DSN()
	log.Printf("🔌 Conectando ao banco de dados: %s", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("❌ Erro ao conectar ao banco: %w", err)
	}

	log.Println("✅ Conexão com o banco estabelecida com sucesso.")

	// Verifica e aplica o schema, se necessário
	if err := ensureSchema(cfg); err != nil {
		log.Printf("⚠️ Erro ao aplicar o schema: %v", err)
	} else {
		log.Println("📄 Verificação e aplicação do schema concluída.")
	}

	return db, nil
}

// ensureSchema verifica se tabelas essenciais existem e aplica o schema.sql se necessário.
func ensureSchema(cfg *config.Config) error {
	sqlDB, err := sql.Open("postgres", cfg.Database.DSN())
	if err != nil {
		return fmt.Errorf("erro ao abrir conexão direta com banco: %w", err)
	}
	defer sqlDB.Close()
	log.Println("📄 verificando se as tabelas existemXXXXXXX.")

	requiredTables := []string{"usuarios", "clinicas"}
	for _, table := range requiredTables {
		exists, err := tableExists(sqlDB, table)
		if err != nil {
			return fmt.Errorf("erro ao verificar existência da tabela '%s': %w", table, err)
		}
		if !exists {
			log.Printf("🚧 Tabela '%s' não encontrada. Aplicando schema.sql...", table)
			return applySchema(sqlDB)
		}
	}

	log.Println("🧩 Todas as tabelas essenciais já existem.")
	return nil
}

// tableExists verifica se uma tabela existe no schema 'public'.
func tableExists(db *sql.DB, tableName string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS (
		SELECT FROM information_schema.tables 
		WHERE table_schema = 'public' AND table_name = $1
	);`
	err := db.QueryRow(query, tableName).Scan(&exists)
	return exists, err
}

// applySchema lê o conteúdo do scripts/schema.sql e executa no banco.
func applySchema(db *sql.DB) error {
	const schemaPath = "scripts/schema.sql"

	script, err := os.ReadFile(schemaPath)
	if err != nil {
		return fmt.Errorf("erro ao ler arquivo %s: %w", schemaPath, err)
	}

	if _, err := db.Exec(string(script)); err != nil {
		return fmt.Errorf("erro ao executar schema.sql: %w", err)
	}

	log.Println("✅ schema.sql executado com sucesso.")
	return nil
}