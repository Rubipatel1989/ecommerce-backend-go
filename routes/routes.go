package routes

import (
	"ecommerce-backend-go/action/admin"
	"ecommerce-backend-go/pkg/sessionx"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/login")
	})

	// Dashboard route (GET)
	app.Get("/dashboard", func(c *fiber.Ctx) error {
		sess, _ := sessionx.Store.Get(c)

		name := sess.Get("admin_name")
		mobile := sess.Get("admin_mobile")
		return c.Render("layout", fiber.Map{
			"Title":  "Admin Dashboard",
			"Name":   name,
			"Mobile": mobile,
		})

	})

	// Login form (GET)
	app.Get("/login", func(c *fiber.Ctx) error {
		errorMsg := c.Query("error", "")
		return c.Render("admin_login", fiber.Map{
			"Error": errorMsg,
		})
	})

	// Login submission (POST)
	app.Post("/login", admin.HandleAdminLogin)

	app.Get("/logout", func(c *fiber.Ctx) error {
		sess, _ := sessionx.Store.Get(c)
		sess.Destroy()
		return c.Redirect("/login")
	})

}
