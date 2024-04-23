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
	IDEscola            int    `json:"id_escola,omitempty"`
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

type Request struct {
	IdGlobal int `json:"id_global"`
	IdUser   int `json:"id_user"`
}

type EscolaResponse struct {
	IdGlobal            int    `json:"id_global"`
	IDEscola            int    `json:"id_escola"`
	Nome                string `json:"nome"`
	CodINEP             string `json:"codinep"`
	Email               string `json:"email"`
	Telefone            string `json:"telefone"`
	NomeCoordenador     string `json:"nome_coordenador"`
	EmailCoordenador    string `json:"email_coordenador"`
	TelefoneCoordenador string `json:"telefone_coordenador"`
	TipoEscola          string `json:"tipo_escola"`
	IDEndereco          int    `json:"idendereco"`
}

type AlunoResponse struct {
	IdGlobal    int      `json:"id_global"`
	IDAluno     int      `json:"id_aluno"`
	IDEscola    int      `json:"id_escola"`
	CPF         string   `json:"cpf"`
	Nome        string   `json:"nome"`
	SerieAtual  string   `json:"serie_atual"`
	Genero      string   `json:"genero"`
	DataNasc    string   `json:"data_nasc"`
	PerfisAcess []string `json:"perfis_acess"`
}
