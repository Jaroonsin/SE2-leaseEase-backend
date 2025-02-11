package middleware

import (
	"LeaseEase/config"
	"LeaseEase/utils"
	"log"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func AuthorizationUserToken(cfg *config.Config) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:     []byte(cfg.JWTApiSecret),
		ErrorHandler:   AuthError,
		SuccessHandler: AuthSuccess,
	})
}

func AuthError(c *fiber.Ctx, err error) error {
	return utils.ErrorResponse(c, fiber.StatusBadRequest, "Unauthorized")
}
func AuthSuccess(c *fiber.Ctx) error {
	return c.Next()
}

func AuthRequired(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		cookie := c.Cookies("auth_token")
		
		if claims, err := utils.ParseJWT(cookie); err != nil {
			return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid token")
		} else {
			// Store token claims in locals if needed for later use in request context
			c.Locals("user", claims)
		}
		log.Println("User claims:", c.Locals("user"))

		return c.Next()
	}
}