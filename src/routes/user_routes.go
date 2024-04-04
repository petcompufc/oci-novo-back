package routes

import (
	"oci-novo/api/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	app.Post("api/<versao>/cadastro", handlers.CreateUser)
}
