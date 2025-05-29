package routes

import (
	adminhandler "ecommerce-backend-go/action/admin"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/admin/login", func(c *fiber.Ctx) error {
		return c.Render("admin_login", fiber.Map{})
	})

	app.Post("/admin/login", adminhandler.HandleAdminLogin)
}
