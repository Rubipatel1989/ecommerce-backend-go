package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// Register each route group
	RegisterAdminRoutes(app)
	RegisterProductRoutes(app)
}
