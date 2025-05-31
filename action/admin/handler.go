package admin

import (
	"ecommerce-backend-go/pkg/sessionx"
	"ecommerce-backend-go/services/responsex"
	"fmt"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

type LoginInput struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func HandleAdminLogin(c *fiber.Ctx) error {
	var input LoginInput
	if err := c.BodyParser(&input); err != nil {
		return responsex.BadRequest(c, "Invalid input")
	}

	name, mobile, err := ValidateAdminCredentials(input.Username, input.Password)
	if err != nil {
		redirectURL := fmt.Sprintf("/login?error=%s", url.QueryEscape(err.Error()))
		return c.Redirect(redirectURL)
	}

	sess, _ := sessionx.Store.Get(c)

	sess.Set("admin_name", name)
	sess.Set("admin_mobile", mobile)
	sess.Save()

	return c.Redirect("/dashboard")
}
