package utils

import (
	"github.com/gofiber/fiber/v2"
)

func SuccessResponse(ctx *fiber.Ctx, status int, message string, data any) error {
	return ctx.Status(status).JSON(fiber.Map{
		"isSuccess": true,
		"message":   message,
		"data":      data,
	})
}

func ErrorResponse(ctx *fiber.Ctx, status int, message string, errors ...interface{}) error {
	resp := fiber.Map{
		"isSuccess": false,
		"message":   message,
	}
	if len(errors) > 0 && errors[0] != nil {
		resp["errors"] = errors[0]
	}
	return ctx.Status(status).JSON(resp)
}

