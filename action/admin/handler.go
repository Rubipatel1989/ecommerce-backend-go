package admin

import (
	"ecommerce-backend-go/services/responsex"

	"github.com/gofiber/fiber/v2"
)

func HandleAdminLogin(c *fiber.Ctx) error {
	type LoginInput struct {
		Username string `form:"username"`
		Password string `form:"password"`
	}

	var input LoginInput
	if err := c.BodyParser(&input); err != nil {
		return responsex.BadRequest(c, "Invalid form input")
	}

	err := ValidateAdminCredentials(input.Username, input.Password)
	if err != nil {
		return responsex.Unauthorized(c, err.Error())
	}

	return responsex.GTSuccess(c, true, "Login successful", nil)
}
