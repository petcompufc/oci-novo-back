package handlers

import (
	"log"
	"oci-novo/api/models"
	"strings"
	"time"

	"github.com/lib/pq"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CreateEscola(c *fiber.Ctx) error {
	// Criar uma instância de DB
	db, err := h.getDB("api_user")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Receba os dados do corpo da solicitação JSON
	var newEscola models.Escola
	if err := c.BodyParser(&newEscola); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Erro ao analisar o corpo de solicitação JSON",
		})
	}

	// Validar os dados de escola
	if newEscola.Cargo != "escola" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Tipo de usuário inválido",
		})
	}

	// Definir formato da senha
	if newEscola.HashSenha == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Senha inválida",
		})
	}

	if newEscola.Nome == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Nome inválido",
		})
	}

	// Validar os dados de endereco

	if newEscola.Endereco.CEP == "" || len(newEscola.Endereco.CEP) != 8 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "CEP inválido",
		})
	}
	if newEscola.Endereco.Bairro == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bairro inválido",
		})
	}
	if newEscola.Endereco.Cidade == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Cidade inválida",
		})
	}
	if newEscola.Endereco.Estado == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Estado inválido",
		})
	}
	if newEscola.Endereco.Rua == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Rua inválida",
		})
	}
	if newEscola.Endereco.Numero <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Número inválido",
		})
	}

	// verificar se o cod_inep já existe na tabela escola
	var codINEPExists string
	err = db.QueryRow("SELECT cod_inep FROM escola WHERE cod_inep = $1", newEscola.CodINEP).Scan(&codINEPExists)
	if err == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "cod_inep já existe",
		})
	}

	// Validar os dados de escola
	// É preciso validar com base em alguma API?
	if newEscola.CodINEP == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Código INEP inválido",
		})
	}

	// Verificar se o email é válido
	if !isValidEmail(newEscola.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Email inválido",
		})
	}

	// Verificar se o email do coordenador é válido
	if !isValidEmail(newEscola.EmailCoordenador) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Email do coordenador inválido",
		})
	}

	// Verificar se o telefone do coordenador é válido
	if newEscola.TelefoneCoordenador == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Telefone do coordenador inválido",
		})
	}

	// Verificar se o tipo de escola é válido
	validTypes := []string{"publica", "particular"}
	if !stringInSlice(newEscola.TipoEscola, validTypes) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Tipo de escola inválido",
		})
	}

	// inserir os dados na tabela usuário
	result := db.QueryRow("INSERT INTO usuario (hash_senha, cargo, ultimo_login) VALUES ($1, $2, $3) RETURNING id_usuario", newEscola.HashSenha, newEscola.Cargo, time.Now())

	// Obter o ID inserido e verificar se houve erro
	var idGlobal int
	err = result.Scan(&idGlobal)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erro ao obter o IdGlobal do usuário inserido",
		})
	}

	// Inserir os dados na tabela endereco
	result = db.QueryRow("INSERT INTO endereco (cep, bairro, cidade, estado, rua, numero, complemento) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id_endereco",
		newEscola.Endereco.CEP,
		newEscola.Endereco.Bairro,
		newEscola.Endereco.Cidade,
		newEscola.Endereco.Estado,
		newEscola.Endereco.Rua,
		newEscola.Endereco.Numero,
		newEscola.Endereco.Complemento,
	)

	// Obter o ID inserido e verificar se houve erro
	var idEndereco int
	err = result.Scan(&idEndereco)
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
		newEscola.CodINEP,
		newEscola.Nome,
		newEscola.Email,
		newEscola.Telefone,
		newEscola.NomeCoordenador,
		newEscola.EmailCoordenador,
		newEscola.TelefoneCoordenador,
		newEscola.TipoEscola,
		idEndereco,
	)
	if err != nil {
		// Pensar em como fazer se usuário for inserido, mas Escola não
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erro ao criar escola",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":    "Usuário criado com sucesso",
		"nome":       newEscola.Nome,
		"id_usuario": idGlobal,
		"cargo":      newEscola.Cargo,
	})
}

func (h *Handler) CreateAluno(c *fiber.Ctx) error {
	// Criar uma instância de DB
	db, err := h.getDB("api_user")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Receba os dados do corpo da solicitação JSON
	var newAluno models.Aluno
	if err := c.BodyParser(&newAluno); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Erro ao analisar o corpo de solicitação JSON",
		})
	}

	// Validar os dados de aluno
	newAluno.Cargo = strings.ToLower(newAluno.Cargo)
	if newAluno.Cargo != "aluno" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Tipo de usuário inválido",
		})
	}

	// Definir formato da senha
	if newAluno.HashSenha == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Senha inválida",
		})
	}

	if newAluno.Nome == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Nome inválido",
		})
	}

	// Verificar se o ID da escola é válido
	_, err = db.Exec("SELECT id_escola FROM escola WHERE id_escola = $1", newAluno.IDEscola)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ID da escola inválido",
		})
	}

	// Verificar se o CPF é válido
	// Verificar formato do CPF
	if len(newAluno.CPF) != 11 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "CPF inválido",
		})
	}

	// Verificar se o CPF já está cadastrado
	var cpf string
	err = db.QueryRow("SELECT cpf FROM aluno WHERE cpf = $1", newAluno.CPF).Scan(&cpf)
	if err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "CPF já cadastrado",
		})
	}

	// Verificar se a série atual é válida
	validSeries := []string{"6f", "7f", "8f", "9f", "1m", "2m", "3m"}
	if !stringInSlice(newAluno.SerieAtual, validSeries) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Série inválida",
		})
	}

	// verificar se o gênero é válido
	validGenders := []string{"masculino", "feminino", "nao-binario", "genero-fluido", "agenero", "bigenero", "transgenero", "outro"}
	if !stringInSlice(newAluno.Genero, validGenders) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Gênero inválido",
		})
	}

	// verificar se a data de nascimento é válida
	if newAluno.DataNasc.IsZero() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Data de nascimento inválida",
		})
	}

	// verificar se os perfis de acessibilidade são válidos
	validPerfis := []string{"fisica", "sensorial", "cognitiva/intelectual", "psiquiatrica", "invisivel"}
	for _, perfil := range newAluno.PerfisAcess {
		if !stringInSlice(perfil, validPerfis) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Perfil de acessibilidade inválido",
			})
		}
	}

	// inserir os dados na tabela usuário
	result := db.QueryRow("INSERT INTO usuario (hash_senha, cargo, ultimo_login) VALUES ($1, $2, $3) RETURNING id_usuario", newAluno.HashSenha, newAluno.Cargo, time.Now())

	// Obter o ID inserido e verificar se houve erro
	var idGlobal int
	err = result.Scan(&idGlobal)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erro ao obter o IdGlobal do usuário inserido",
		})
	}

	// Inserir os dados na tabela aluno
	_, err = db.Exec("INSERT INTO aluno (id_global, id_escola, cpf, nome, serie_atual, genero, data_nasc, perfis_acess) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		idGlobal,
		newAluno.IDEscola,
		newAluno.CPF,
		newAluno.Nome,
		newAluno.SerieAtual,
		newAluno.Genero,
		newAluno.DataNasc,
		pq.Array(newAluno.PerfisAcess),
	)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erro ao criar aluno",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":    "Aluno criado com sucesso",
		"nome":       newAluno.Nome,
		"id_usuario": idGlobal,
		"cargo":      newAluno.Cargo,
	})
}

func (h *Handler) GetEscolas(c *fiber.Ctx) error {
	request := models.Request{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Erro ao analisar o corpo de solicitação JSON",
		})
	}

	// Verificar se os IDs são válidos
	if request.IdGlobal <= 0 || request.IdUser <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ID inválido",
		})
	}

	// Criar instância de DB
	db, err := h.getDB("api_user")
	if err != nil {
		log.Fatal(err)
	}

	row := db.QueryRow("SELECT cargo FROM usuario WHERE id_usuario = $1", request.IdGlobal)

	var cargo string
	err = row.Scan(&cargo)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erro ao obter cargo do usuário",
		})
	}
	db.Close()

	switch cargo {
	case "escola":

		// Criar instância de DB
		db, err := h.getDB("escola_user")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		row := db.QueryRow("SELECT * FROM escola WHERE id_escola = $1", request.IdUser)

		escola := models.EscolaResponse{}
		err = row.Scan(&escola.IDEscola, &escola.IdGlobal, &escola.CodINEP, &escola.Nome, &escola.IDEndereco, &escola.Email, &escola.Telefone, &escola.NomeCoordenador, &escola.EmailCoordenador, &escola.TelefoneCoordenador, &escola.TipoEscola)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Erro ao obter escola",
			})
		}

		// Se o idglobal da requisição (no caso de ser uma escola fazendo a requisição)
		// for diferente do idglobal da escola que está sendo requisitada, então
		// a escola não tem permissão para acessar esses dados
		if escola.IdGlobal != request.IdGlobal {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Você não tem permissão para acessar esses dados",
			})
		}

		// Caso contrário, vamos buscar o endereço da escola
		row = db.QueryRow("SELECT cep, bairro, cidade, estado, rua, numero, complemento FROM endereco WHERE id_endereco = $1", escola.IDEndereco)
		err = row.Scan(&escola.Endereco.CEP, &escola.Endereco.Bairro, &escola.Endereco.Cidade, &escola.Endereco.Estado, &escola.Endereco.Rua, &escola.Endereco.Numero, &escola.Endereco.Complemento)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Erro ao obter endereço da escola",
			})
		}

		return c.Status(fiber.StatusOK).JSON(escola)

	case "petiano":
		// Criar instância de DB
		db, err := h.getDB("petiano_user")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		row := db.QueryRow("SELECT * FROM escola WHERE id_escola = $1", request.IdUser)

		escola := models.EscolaResponse{}
		err = row.Scan(&escola.IDEscola, &escola.IdGlobal, &escola.CodINEP, &escola.Nome, &escola.IDEndereco, &escola.Email, &escola.Telefone, &escola.NomeCoordenador, &escola.EmailCoordenador, &escola.TelefoneCoordenador, &escola.TipoEscola)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Erro ao obter escola",
			})
		}

		// Buscar o endereço da escola
		row = db.QueryRow("SELECT cep, bairro, cidade, estado, rua, numero, complemento FROM endereco WHERE id_endereco = $1", escola.IDEndereco)
		err = row.Scan(&escola.Endereco.CEP, &escola.Endereco.Bairro, &escola.Endereco.Cidade, &escola.Endereco.Estado, &escola.Endereco.Rua, &escola.Endereco.Numero, &escola.Endereco.Complemento)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Erro ao obter endereço da escola",
			})
		}

		return c.Status(fiber.StatusOK).JSON(escola)

	default:
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Você não tem permissão para acessar esses dados",
		})
	}
}

func (h *Handler) GetAluno(c *fiber.Ctx) error {
	request := models.Request{}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Erro ao analisar o corpo de solicitação JSON",
		})
	}

	// Criar instância de DB
	db, err := h.getDB("api_user")
	if err != nil {
		log.Fatal(err)
	}

	row := db.QueryRow("SELECT cargo FROM usuario WHERE id_usuario = $1", request.IdGlobal)
	var cargo string
	err = row.Scan(&cargo)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erro ao obter cargo do usuário",
		})
	}

	db.Close()

	switch cargo {
	case "escola":
		// Criar instância de DB
		db, err := h.getDB("escola_user")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		aluno := models.AlunoResponse{}
		var perfisAccess string
		row = db.QueryRow("SELECT * FROM aluno WHERE id_aluno = $1", request.IdUser)
		err = row.Scan(&aluno.IDAluno, &aluno.IdGlobal, &aluno.IDEscola, &aluno.CPF, &aluno.Nome, &aluno.SerieAtual, &aluno.Genero, &aluno.DataNasc, &perfisAccess)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Erro ao obter aluno",
			})
		}
		aluno.PerfisAcess = strings.Split(perfisAccess, ",")

		var idGlobalEscola int
		row = db.QueryRow("SELECT id_global FROM escola WHERE id_escola = $1", aluno.IDEscola)
		err = row.Scan(&idGlobalEscola)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Erro ao obter id_global da escola",
			})
		}

		if idGlobalEscola != request.IdGlobal {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Você não tem permissão para acessar esses dados",
			})
		}

		return c.Status(fiber.StatusOK).JSON(aluno)

	case "petiano":
		// Criar instância de DB
		db, err := h.getDB("petiano_user")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		aluno := models.AlunoResponse{}
		var perfisAccess string
		row = db.QueryRow("SELECT * FROM aluno WHERE id_aluno = $1", request.IdUser)
		err = row.Scan(&aluno.IDAluno, &aluno.IdGlobal, &aluno.IDEscola, &aluno.CPF, &aluno.Nome, &aluno.SerieAtual, &aluno.Genero, &aluno.DataNasc, &perfisAccess)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Erro ao obter aluno",
			})
		}
		aluno.PerfisAcess = strings.Split(perfisAccess, ",")

		return c.Status(fiber.StatusOK).JSON(aluno)

	case "aluno":
		if request.IdGlobal != request.IdUser {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Você somente pode acessar seus próprios dados",
			})
		}

		// Criar instância de DB
		db, err := h.getDB("aluno_user")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		aluno := models.AlunoResponse{}
		var perfisAccess string
		row = db.QueryRow("SELECT * FROM aluno WHERE id_aluno = $1", request.IdUser)
		err = row.Scan(&aluno.IDAluno, &aluno.IdGlobal, &aluno.IDEscola, &aluno.CPF, &aluno.Nome, &aluno.SerieAtual, &aluno.Genero, &aluno.DataNasc, &perfisAccess)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Erro ao obter aluno",
			})
		}
		aluno.PerfisAcess = strings.Split(perfisAccess, ",")

		if aluno.IdGlobal != request.IdGlobal {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Você só tem permisão para acessar seus próprios dados",
			})
		}

		return c.Status(fiber.StatusOK).JSON(aluno)

	default:
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Você não tem permissão para acessar esses dados",
		})
	}
}
