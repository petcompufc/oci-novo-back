package handlers

import (
	"database/sql"
	"oci-novo/api/models"
	"time"

	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	db, err := sql.Open("postgres", "postgres://api_user:1234@localhost:5432/oci_dados?sslmode=disable")
	if err != nil {
		return err
	}

	defer db.Close()

	// Receba os dados do corpo da solicitação JSON
	var newUser models.Usuario
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Erro ao analisar o corpo de solicitação JSON",
		})
	}

	// Validação dos dados do usuário
	// (Necessário melhorar a validação, este é apenas um exemplo)
	if newUser.IdGlobal <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "IDGlobal inválido",
		})
	}

	// recebendo o cargo da solicitação JSON
	cargo, ok := newUser.Dados["cargo"].(string)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Campo 'cargo' inválido",
		})
	}

	// inserir os dados na tabela usuário
	_, err = db.Exec("INSERT INTO usuario (hash_senha, cargo, ultimo_login) VALUES ($1, $2, $3)", newUser.Dados["hash_senha"], cargo, time.Now())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erro ao criar usuário",
		})
	}

	// Obter o último ID inserido
	var idGlobal int
	err = db.QueryRow("SELECT lastval()").Scan(&idGlobal)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erro ao obter o último ID inserido",
		})
	}

	if cargo == "aluno" {
		// Inserir os dados na tabela aluno
		_, err = db.Exec("INSERT INTO aluno (id_global, id_escola, cpf, nome, serie_atual, genero, data_nasc, perfis_aces) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
			idGlobal,
			newUser.Dados["id_escola"],
			newUser.Dados["cpf"],
			newUser.Dados["nome"],
			newUser.Dados["serie_atual"],
			newUser.Dados["genero"],
			newUser.Dados["data_nasc"],
			newUser.Dados["perfis_aces"])
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Erro ao criar aluno",
			})
		}

	} else if cargo == "escola" {
		//
	}
}
