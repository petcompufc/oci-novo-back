package routes

import (
	"oci-novo/api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App, h *handlers.Handlers) {
	app.Post("api/<versao>/cadastro", h.CreateUser)
}
