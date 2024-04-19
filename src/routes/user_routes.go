package routes

import (
	"oci-novo/api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App, handlers *handlers.Handler) {
	// alterar endpoint para seguir padrão "api/<versao>/cadastro"

	app.Post("api/v1/cadastro/escola", handlers.CreateEscola)

	app.Post("api/v1/cadastro/aluno", handlers.CreateAluno)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
