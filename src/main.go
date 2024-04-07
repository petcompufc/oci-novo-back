package main

import (
	"database/sql"
	"log"
	"oci-novo/api/handlers"
	"oci-novo/api/routes"

	"strconv"

	"fmt"

	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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

	h := &handlers.Handlers{DB: db}
	app := fiber.New()
	routes.SetupUserRoutes(app, h)
	app.Listen(":" + env_port)

}
