package middleware

import (
	"LeaseEase/config"
	"LeaseEase/utils"

	"github.com/gofiber/fiber/v2"
)

func AuthRequired(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		cookie := c.Cookies("auth_token")
		
		if claims, err := utils.ParseJWT(cookie); err != nil {
			return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid token")
		} else {
			// Store token claims in locals if needed for later use in request context
			c.Locals("user", claims)
		}
		

		return c.Next()
	}
}