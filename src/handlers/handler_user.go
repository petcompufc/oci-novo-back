package handlers

import (
	"database/sql"
	"oci-novo/api/models"
	"time"

	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2"
)

type Handlers struct {
	DB *sql.DB
}

func (h *Handlers) CreateUser(c *fiber.Ctx) error {
	// Receba os dados do corpo da solicitação JSON
	var newUser models.Usuario
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Erro ao analisar o corpo de solicitação JSON",
		})
	}

	// Validação dos dados do usuário
	// ...

	// recebendo o cargo da solicitação JSON
	cargo := newUser.Cargo

	// inserir os dados na tabela usuário
	// inserir os dados na tabela usuário
	result, err := h.DB.Exec("INSERT INTO usuario (hash_senha, cargo, ultimo_login) VALUES ($1, $2, $3) RETURNING id_usuario", newUser.HashSenha, newUser.Cargo, time.Now())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erro ao criar usuário",
		})
	}

	// Obter o ID inserido
	id, err := result.LastInsertId()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erro ao obter o ID inserido",
		})
	}
	idGlobal := int(id)

	if cargo == "aluno" {
		// Inserir os dados na tabela aluno
		_, err = h.DB.Exec("INSERT INTO aluno (id_global, id_escola, cpf, nome, serie_atual, genero, data_nasc, perfis_acess) VALUES ($1, $2, $3, $4, $5, $6, $7)",
			idGlobal,
			newUser.IDEscola,
			newUser.CPF,
			newUser.Nome,
			newUser.SerieAtual,
			newUser.Genero,
			newUser.DataNasc,
			newUser.PerfisAcess,
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Erro ao criar aluno",
			})
		}

	} else if cargo == "escola" {
		// Inserir os dados na tabela endereco
		result, err := h.DB.Exec("INSERT INTO endereco (cep, bairro, cidade, estado, rua, numero, complemento) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id_endereco",
			newUser.CEP,
			newUser.Bairro,
			newUser.Cidade,
			newUser.Estado,
			newUser.Rua,
			newUser.Numero,
			newUser.Complemento,
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Erro ao criar endereço",
			})
		}

		// Obter o ID inserido
		idEndereco, err := result.LastInsertId()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Erro ao obter o ID inserido",
			})
		}

		// Inserir os dados na tabela escola
		_, err = h.DB.Exec("INSERT INTO escola (id_global, cod_inep, nome, email, telefone, nome_coordenador, email_coordenador, telefone_coordenador, tipo_escola, id_endereco) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
			idGlobal,
			newUser.CodINEP,
			newUser.Nome,
			newUser.Email,
			newUser.Telefone,
			newUser.NomeCoordenador,
			newUser.EmailCoordenador,
			newUser.TelefoneCoordenador,
			newUser.TipoEscola,
			idEndereco,
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Erro ao criar escola",
			})
		}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":    "Usuário criado com sucesso",
		"id_usuario": idGlobal,
		"cargo":      cargo,
	})
}