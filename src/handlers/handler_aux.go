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
	if db, ok := h.dbs[user]; ok {
		return db, nil
	}

	password := os.Getenv(strings.ToUpper(user) + "_PWD")
	connectionString := fmt.Sprintf("postgres://%s:%s@localhost:5432/oci_dados?sslmode=disable", user, password)
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
	if !stringInSlice("@", []string{email}) {
		return false
	}

	if !stringInSlice(".", []string{email}) {
		return false
	}

	return true
}
