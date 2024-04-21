package models

import "time"

type Aluno struct {
	IdGlobal    int       `json:"id_global"`
	HashSenha   string    `json:"hash_senha"`
	Cargo       string    `json:"cargo"`
	UltimoLogin time.Time `json:"ultimo_login"`

	// Campos da estrutura Aluno
	IDAluno     int       `json:"id_aluno,omitempty"`
	IDEscola    int       `json:"id_escola,omitempty"`
	CPF         string    `json:"cpf,omitempty"`
	Nome        string    `json:"nome,omitempty"`
	SerieAtual  string    `json:"serie_atual,omitempty"`
	Genero      string    `json:"genero,omitempty"`
	DataNasc    time.Time `json:"data_nasc,omitempty"`
	PerfisAcess []string  `json:"perfis_acess,omitempty"`
}

type Endereco struct {
	CEP         string `json:"cep,omitempty"`
	Bairro      string `json:"bairro,omitempty"`
	Cidade      string `json:"cidade,omitempty"`
	Estado      string `json:"estado,omitempty"`
	Rua         string `json:"rua,omitempty"`
	Numero      int    `json:"numero,omitempty"`
	Complemento string `json:"complemento,omitempty"`
}

type Escola struct {
	IdGlobal    int       `json:"id_global"`
	HashSenha   string    `json:"hash_senha"`
	Cargo       string    `json:"cargo"`
	UltimoLogin time.Time `json:"ultimo_login"`

	// Campos da estrutura Escola
	Nome                string `json:"nome,omitempty"`
	CodINEP             string `json:"codinep,omitempty"`
	Email               string `json:"email,omitempty"`
	Telefone            string `json:"telefone,omitempty"`
	NomeCoordenador     string `json:"nome_coordenador,omitempty"`
	EmailCoordenador    string `json:"email_coordenador,omitempty"`
	TelefoneCoordenador string `json:"telefone_coordenador,omitempty"`
	TipoEscola          string `json:"tipo_escola,omitempty"`
	IDEndereco          int    `json:"idendereco,omitempty"`

	Endereco Endereco `json:"endereco,omitempty"`
}
