package routes

import (
	"ecommerce-backend-go/action/product"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/product", product.AddProduct)
}
