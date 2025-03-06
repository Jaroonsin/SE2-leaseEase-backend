package utils

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	StatusCode int         `json:"status_code" example:"888"`
	Message    string      `json:"message" example:"message"`
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
