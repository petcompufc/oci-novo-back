package main

import (
	"log"
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

	app := fiber.New()
	routes.SetupUserRoutes(app)
	app.Listen(":" + env_port)

}
