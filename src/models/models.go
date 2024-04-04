package models

import (
	"time"
)

type Usuario struct {
	IdGlobal int                    `json:"id_global"`
	Dados    map[string]interface{} `json:"dados"`
}

type Escola struct {
	HashSenha           string    `json:"-"`
	Cargo               string    `json:"cargo"`
	UltimoLogin         time.Time `json:"ultimo_login"`
	IDEscola            int       `json:"-"`
	CodINEP             string    `json:"-"`
	Nome                string    `json:"-"`
	Email               string    `json:"-"`
	Telefone            string    `json:"-"`
	NomeCoordenador     string    `json:"-"`
	EmailCoordenador    string    `json:"-"`
	TelefoneCoordenador string    `json:"-"`
	TipoEscola          string    `json:"-"`
	IDEndereco          int       `json:"-"`
	Endereco            *Endereco `json:"-"`
}

type Endereco struct {
	ID          int    `json:"id"`
	CEP         string `json:"cep"`
	Bairro      string `json:"bairro"`
	Cidade      string `json:"cidade"`
	Estado      string `json:"estado"`
	Rua         string `json:"rua"`
	Numero      int    `json:"numero"`
	Complemento string `json:"complemento"`
}

type Aluno struct {
	HashSenha   string    `json:"-"`
	Cargo       string    `json:"cargo"`
	UltimoLogin time.Time `json:"ultimo_login"`
	IDAluno     int       `json:"-"`
	IDEscola    int       `json:"-"`
	CPF         string    `json:"-"`
	Nome        string    `json:"-"`
	SerieAtual  string    `json:"-"`
	Genero      string    `json:"-"`
	DataNasc    string    `json:"-"`
	PerfisAcess []string  `json:"-"`
}

func CreateUser(user *Usuario) error {

}
