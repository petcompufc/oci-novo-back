package main

import (
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

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%s", env_port)))
}
