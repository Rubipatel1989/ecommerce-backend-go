package routes

import "github.com/gofiber/fiber/v2"

func RegisterProductRoutes(app *fiber.App) {
	app.Get("/products", func(c *fiber.Ctx) error {
		return c.SendString("Product list")
	})

	app.Get("/products/add", func(c *fiber.Ctx) error {
		return c.SendString("Add product form")
	})

	app.Post("/products/add", func(c *fiber.Ctx) error {
		return c.SendString("Handle add product")
	})

	app.Get("/products/edit/:id", func(c *fiber.Ctx) error {
		return c.SendString("Edit product with ID: " + c.Params("id"))
	})

	app.Post("/products/edit/:id", func(c *fiber.Ctx) error {
		return c.SendString("Update product with ID: " + c.Params("id"))
	})

	app.Get("/products/delete/:id", func(c *fiber.Ctx) error {
		return c.SendString("Delete product with ID: " + c.Params("id"))
	})
}
