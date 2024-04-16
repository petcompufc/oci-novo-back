package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"oci-novo/api/models"
	"os"
	"time"

	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	// Criar uma instância de DB
	password := os.Getenv("API_USER_PWD")
	connectionString := fmt.Sprintf("postgres://api_user:%s@localhost:5432/oci_dados?sslmode=disable", password)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Receba os dados do corpo da solicitação JSON
	var newUser models.Usuario
	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Erro ao analisar o corpo de solicitação JSON",
		})
	}

	// Validar os dados de usuário
	if newUser.HashSenha == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Senha inválida",
		})
	}

	if newUser.Nome == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Nome inválido",
		})
	}

	//  Verificar o tipo do usuário e fazer a validação apropriada
	switch newUser.Cargo {
	case "aluno":
		// Validar os dados de aluno

		// Verificar se o ID da escola é válido
		_, err := db.Exec("SELECT id_escola FROM escola WHERE id_escola = $1", newUser.IDEscola)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "ID da escola inválido",
			})
		}

		// Verificar se o CPF é válido
		// Verificar formato do CPF
		if len(newUser.CPF) != 11 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "CPF inválido",
			})
		}

		// Verificar se o CPF já está cadastrado
		var cpf string
		err = db.QueryRow("SELECT cpf FROM aluno WHERE cpf = $1", newUser.CPF).Scan(&cpf)
		if err == nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "CPF já cadastrado",
			})
		}

		// Verificar se a série atual é válida
		validSeries := []string{"6f", "7f", "8f", "9f", "1m", "2m", "3m"}
		if !stringInSlice(newUser.SerieAtual, validSeries) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Série inválida",
			})
		}

		// verificar se o gênero é válido
		validGenders := []string{"masculino", "feminino", "nao-binario", "genero-fluido", "agenero", "bigenero", "transgenero", "outro"}
		if !stringInSlice(newUser.Genero, validGenders) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Gênero inválido",
			})
		}

		// verificar se a data de nascimento é válida
		if newUser.DataNasc.IsZero() {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Data de nascimento inválida",
			})
		}

		// verificar se os perfis de acessibilidade são válidos
		validPerfis := []string{"fisica", "sensorial", "cognitiva/intelectual", "psiquiatrica", "invisivel"}
		for _, perfil := range newUser.PerfisAcess {
			if !stringInSlice(perfil, validPerfis) {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": "Perfil de acessibilidade inválido",
				})
			}
		}

	case "escola":

		// Validar os dados de endereco

		if newUser.CEP == "" || len(newUser.CEP) != 8 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "CEP inválido",
			})
		}
		if newUser.Bairro == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Bairro inválido",
			})
		}
		if newUser.Cidade == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Cidade inválida",
			})
		}
		if newUser.Estado == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Estado inválido",
			})
		}
		if newUser.Rua == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Rua inválida",
			})
		}
		if newUser.Numero <= 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Número inválido",
			})
		}

		// Verificar a partir de consultas
		// verificar se o cod_inep já existe na tabela escola
		var codINEPExists string
		err = db.QueryRow("SELECT cod_inep FROM escola WHERE cod_inep = $1", newUser.CodINEP).Scan(&codINEPExists)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Erro ao verificar a existência do cod_inep",
			})
		}
		if codINEPExists != "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "cod_inep já existe",
			})
		}

		// Validar os dados de escola
		// É preciso validar com base em alguma API?
		if newUser.CodINEP == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Código INEP inválido",
			})
		}

		// Verificar se o email é válido
		if !isValidEmail(newUser.Email) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Email inválido",
			})
		}

		// Verificar se o email do coordenador é válido
		if !isValidEmail(newUser.EmailCoordenador) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Email do coordenador inválido",
			})
		}

		// Verificar se o telefone do coordenador é válido
		if newUser.TelefoneCoordenador != "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Telefone do coordenador inválido",
			})
		}

		// Verificar se o tipo de escola é válido
		validTypes := []string{"publica", "particular"}
		if !stringInSlice(newUser.TipoEscola, validTypes) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Tipo de escola inválido",
			})
		}

	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Tipo de usuário inválido",
		})
	}

	// inserir os dados na tabela usuário
	result := db.QueryRow("INSERT INTO usuario (hash_senha, cargo, ultimo_login) VALUES ($1, $2, $3) RETURNING id_usuario", newUser.HashSenha, newUser.Cargo, time.Now())

	// Obter o ID inserido e verificar se houve erro
	var idGlobal int
	err = result.Scan(&idGlobal)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erro ao obter o IdGlobal do usuário inserido",
		})
	}

	// Verificar o tipo do usuário e inserir os dados apropriados
	switch newUser.Cargo {
	case "aluno":
		// Inserir os dados na tabela aluno
		_, err = db.Exec("INSERT INTO aluno (id_global, id_escola, cpf, nome, serie_atual, genero, data_nasc, perfis_acess) VALUES ($1, $2, $3, $4, $5, $6, $7)",
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

	case "escola":
		// Inserir os dados na tabela endereco
		result := db.QueryRow("INSERT INTO endereco (cep, bairro, cidade, estado, rua, numero, complemento) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id_endereco",
			newUser.CEP,
			newUser.Bairro,
			newUser.Cidade,
			newUser.Estado,
			newUser.Rua,
			newUser.Numero,
			newUser.Complemento,
		)

		// Obter o ID inserido e verificar se houve erro
		var idEndereco int
		err := result.Scan(&idEndereco)
		if err != nil {
			// Pensar em como fazer se usuário for inserido, mas Endereço não
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Erro ao obter o IdEndereco do endereço inserido",
			})
		}

		// Inserir os dados na tabela escola
		_, err = db.Exec("INSERT INTO escola (id_global, cod_inep, nome, email, telefone, nome_coordenador, email_coordenador, telefone_coordenador, tipo_escola, id_endereco) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
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
			// Pensar em como fazer se usuário for inserido, mas Escola não
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Erro ao criar escola",
			})
		}

	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Tipo de usuário inválido",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":    "Usuário criado com sucesso",
		"id_usuario": idGlobal,
		"cargo":      newUser.Cargo,
	})
}

// Função auxiliar para verificar se uma string está em uma lista de strings
func stringInSlice(a string, list []string) bool {
	for _, v := range list {
		if v == a {
			return true
		}
	}
	return false
}

func isValidEmail(email string) bool {
	// verificar se string email contém @
	if !stringInSlice("@", []string{email}) {
		return false
	}

	if !stringInSlice(".", []string{email}) {
		return false
	}

	return true
}
