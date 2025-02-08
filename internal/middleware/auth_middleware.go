package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func AuthorizationUserToken() fiber.Handler {
    return jwtware.New(jwtware.Config{
        SigningKey:   []byte(os.Getenv("JWT_SECRET")),
        ErrorHandler: AuthError, 
        SuccessHandler: AuthSuccess, 
    })
}


func AuthError(c *fiber.Ctx, err error) error {
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error":   "Unauthorized",
			"message": err.Error(),
		})
	}
	return c.Next()
}
func AuthSuccess(c *fiber.Ctx) error { 
    return c.Next()              
}