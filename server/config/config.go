package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Config armazena todas as configurações da aplicação
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

// ServerConfig armazena configurações do servidor HTTP
type ServerConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

// DatabaseConfig armazena configurações do banco de dados
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// JWTConfig armazena configurações do JWT
type JWTConfig struct {
	Secret          string
	AccessTokenExp  time.Duration
	RefreshTokenExp time.Duration
}

// Load carrega as configurações do ambiente
func Load() (*Config, error) {
	// Carregar variáveis de ambiente do arquivo .env se existir
	_ = godotenv.Load()

	// Configurações do servidor
	port := getEnv("SERVER_PORT", "4000")
	readTimeout, _ := strconv.Atoi(getEnv("SERVER_READ_TIMEOUT", "10"))
	writeTimeout, _ := strconv.Atoi(getEnv("SERVER_WRITE_TIMEOUT", "10"))
	idleTimeout, _ := strconv.Atoi(getEnv("SERVER_IDLE_TIMEOUT", "60"))

	// Configurações do banco de dados
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "erp_system")
	dbSSLMode := getEnv("DB_SSLMODE", "disable")

	// Configurações do JWT
	jwtSecret := getEnv("JWT_SECRET", "your-secret-key")
	jwtAccessExp, _ := strconv.Atoi(getEnv("JWT_ACCESS_EXP", "15"))      // 15 minutos
	jwtRefreshExp, _ := strconv.Atoi(getEnv("JWT_REFRESH_EXP", "10080")) // 7 dias

	return &Config{
		Server: ServerConfig{
			Port:         port,
			ReadTimeout:  time.Duration(readTimeout) * time.Second,
			WriteTimeout: time.Duration(writeTimeout) * time.Second,
			IdleTimeout:  time.Duration(idleTimeout) * time.Second,
		},
		Database: DatabaseConfig{
			Host:     dbHost,
			Port:     dbPort,
			User:     dbUser,
			Password: dbPassword,
			DBName:   dbName,
			SSLMode:  dbSSLMode,
		},
		JWT: JWTConfig{
			Secret:          jwtSecret,
			AccessTokenExp:  time.Duration(jwtAccessExp) * time.Minute,
			RefreshTokenExp: time.Duration(jwtRefreshExp) * time.Minute,
		},
	}, nil
}

// DSN retorna a string de conexão com o banco de dados
func (c *DatabaseConfig) DSN() string {
	fmt.Printf("\nhost=%s port=%s user=%s password=%s dbname=%s sslmode=%s\n",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode)

	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode,
	)
}

// getEnv retorna o valor da variável de ambiente ou o valor padrão
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
