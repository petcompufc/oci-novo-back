package routes

import (
	"oci-novo/api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App, h *handlers.Handlers) {
	// alterar endpoint para seguir padrão "api/<versao>/cadastro"
	app.Post("api/1/cadastro", h.CreateUser)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
