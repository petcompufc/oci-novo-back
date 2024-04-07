package main

import (
	"database/sql"
	"log"

	"strconv"

	"fmt"

	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	env_port := os.Getenv("API_PORT")
	port, err := strconv.Atoi(env_port)

	if err != nil || port < 0 {
		fmt.Printf("'%s' is a invalid port number.\n", env_port)
		os.Exit(2)
	}

	// Criar uma instância de DB
	connectionString := "postgres://api_user:1234@localhost:5432/oci_dados?sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%s", env_port)))
}
