package handlers

import (
	"context"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/msd/server/database"
	"github.com/msd/server/models"
	"golang.org/x/crypto/bcrypt"
)

func CriarUsuario(c *fiber.Ctx) error {
	var usuario models.Usuario

	if err := c.BodyParser(&usuario); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"erro": "Dados inválidos"})
	}

	if usuario.Email == "" || usuario.Senha == "" || usuario.Nome == "" || usuario.Perfil == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"erro": "Campos obrigatórios não informados"})
	}

	// Criptografar senha
	hash, err := bcrypt.GenerateFromPassword([]byte(usuario.Senha), 10)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"erro": "Erro ao criptografar senha"})
	}

	usuario.Senha = string(hash)

	query := `
		INSERT INTO usuarios (	
								nome, 
								email, 
								senha, 
								perfil, 
								clinica_id, 
								supervisor_id, 
								ativo, 
								data_inicio_inatividade, 
								data_fim_inatividade, 
								motivo_inatividade
							)
		VALUES ($1, $2, $3, $4, $5, $6, TRUE, $7, $8, $9)
		RETURNING id
	`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = database
			.DB
			.QueryRow(	ctx, 
						query,
							usuario.Nome, 
							usuario.Email, 
							usuario.Senha, 
							usuario.Perfil,
							usuario.ClinicaID, 
							usuario.SupervisorID,
							usuario.DataInicioInatividade, 
							usuario.DataFimInatividade, 
							usuario.MotivoInatividade,
	).Scan(&usuario.ID)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"erro": "Erro ao cadastrar usuário"})
	}

	usuario.Senha = "" // Não retornar a senha
	return c.Status(fiber.StatusCreated).JSON(usuario)
}
