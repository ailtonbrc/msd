# Simple ERP Service

Backend em Golang para um sistema ERP simples, utilizando Gin e GORM.

## Requisitos

- Go 1.21 ou superior
- PostgreSQL 12 ou superior

## Configuração

1. Clone o repositório
2. Configure o arquivo `.env` com suas credenciais de banco de dados
3. Execute o script SQL para criar o banco de dados e tabelas (disponível em `migrations/create_database.sql`) Comando: `psql -U postgres -f migrations/create_database.sql`
4. Execute o comando `go mod tidy` para instalar as dependências
5. Execute o comando `go run cmd/api/main.go` para iniciar o servidor

## Estrutura do Projeto
