package admin

import (
	"github.com/gofiber/fiber/v2"
)

type LoginInput struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func HandleAdminLogin(c *fiber.Ctx) error {
	var input LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid input")
	}

	err := ValidateAdminCredentials(input.Username, input.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
	}

	return c.Redirect("/admin/dashboard")
}
