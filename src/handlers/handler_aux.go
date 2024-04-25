package handlers

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
)

// Handler é uma estrutura que contém um mapa de conexões de banco de dados
type Handler struct {
	dbs map[string]*sql.DB
}

// Funções auxiliares para manipulação de banco de dados

// NewHandler cria uma nova instância de Handler
func NewHandler() *Handler {
	return &Handler{dbs: make(map[string]*sql.DB)}
}

func (h *Handler) getDB(user string) (*sql.DB, error) {
	password := os.Getenv(strings.ToUpper(user) + "_PWD")
	connectionString := fmt.Sprintf("postgres://%s:%s@localhost:5434/oci_dados?sslmode=disable", user, password)

	if db, ok := h.dbs[user]; ok {
		// Verificar se a conexão com o banco de dados ainda está ativa
		if err := db.Ping(); err != nil {
			// Se a conexão estiver fechada, reabra-a.
			connectionString := fmt.Sprintf("postgres://%s:1234@localhost:5434/oci_dados?sslmode=disable", user)
			db, err = sql.Open("postgres", connectionString)
			if err != nil {
				return nil, err
			}
		}
		return db, nil
	}

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	h.dbs[user] = db
	return db, nil
}

// Funções auxiliares para manipulação de dados

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
	if !stringInSlice("@", strings.Split(email, "")) {
		return false
	}

	if !stringInSlice(".", strings.Split(email, "")) {
		return false
	}

	return true
}
