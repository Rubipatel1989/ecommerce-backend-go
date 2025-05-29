package product

import (
	"ecommerce-backend-go/services/responsex"
	"ecommerce-backend-go/services/validation"

	"github.com/gofiber/fiber/v2"
)

type ProductInput struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

type ProductResponse struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

type ProductListResponse struct {
	Products []ProductResponse `json:"products"`
}

func AddProduct(c *fiber.Ctx) error {
	var input validation.ProductInput
	if err := c.BodyParser(&input); err != nil {
		return responsex.BadRequest(c, "Invalid request")
	}

	if err := input.Validate(); err != nil {
		return responsex.BadRequest(c, err.Error())
	}

	err := CreateProduct(input.Name, input.Price, input.Stock)
	if err != nil {
		return responsex.GTError(c, 500, "Failed to insert product")
	}

	return responsex.GTSuccess(c, true, "Product added successfully", nil)
}
