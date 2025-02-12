package utils

import (
	"github.com/gofiber/fiber/v2"
)

// Account model info
// @Description User account information
// @Description with user id and username
type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func SuccessResponse(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	return c.Status(statusCode).JSON(Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	})
}

func ErrorResponse(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(Response{
		StatusCode: statusCode,
		Message:    message,
	})
}
