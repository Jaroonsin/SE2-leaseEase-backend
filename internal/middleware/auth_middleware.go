package middleware

import (
	"LeaseEase/config"
	"LeaseEase/utils"

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
