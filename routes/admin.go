package routes

import (
	"ecommerce-backend-go/action/admin"
	"ecommerce-backend-go/pkg/sessionx"

	"github.com/gofiber/fiber/v2"
)

func RegisterAdminRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/login")
	})

	app.Get("/dashboard", func(c *fiber.Ctx) error {
		sess, _ := sessionx.Store.Get(c)
		return c.Render("layout", fiber.Map{
			"Title":  "Admin Dashboard",
			"Name":   sess.Get("admin_name"),
			"Mobile": sess.Get("admin_mobile"),
		})
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		errorMsg := c.Query("error", "")
		return c.Render("admin/login", fiber.Map{
			"Error": errorMsg,
		})
	})

	app.Post("/login", admin.HandleAdminLogin)

	app.Get("/logout", func(c *fiber.Ctx) error {
		sess, _ := sessionx.Store.Get(c)
		sess.Destroy()
		return c.Redirect("/login")
	})
}
