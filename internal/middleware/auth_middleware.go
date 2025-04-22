package middleware

import (
	"LeaseEase/config"
	"LeaseEase/utils"
	"log"

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
		log.Println("Auth middleware passed")
		return c.Next()
	}
}

func AdminRoleRequired(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		cookie := c.Cookies("auth_token")
		claims, err := utils.ParseJWT(cookie)
		if err != nil {
			return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid token")
		}
		if claims["role"] != "admin" {
			return utils.ErrorResponse(c, fiber.StatusForbidden, "Forbidden")
		} else {
			log.Println("Admin role middleware passed")
			return c.Next()
		}
	}
}
