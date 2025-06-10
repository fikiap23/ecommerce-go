package utils

import (
	"fmt"
	"go-ecommerce-app/pkg/errors"
	"go-ecommerce-app/pkg/locales"

	"github.com/gofiber/fiber/v2"
)


type CustomError struct {
	Message    string
	StatusCode int
	Detail     any
	Key     errors.ErrorKey
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("%s (code: %d)", e.Message, e.StatusCode)
}

func NewCustomError(key errors.ErrorKey, statusCode int, lang locales.Language, detail ...any) *CustomError {
	var errDetail any
	if len(detail) > 0 {
		errDetail = detail[0]
	}
	return &CustomError{
		Key:     key,
		Message: locales.GetMessage(key, lang),
		StatusCode: statusCode,
		Detail:     errDetail,
	}
}

func HandleError(ctx *fiber.Ctx, err error) error {
	if customErr, ok := err.(*CustomError); ok {
		return ErrorResponse(ctx, customErr.StatusCode, customErr.Message, customErr.Detail)
	}
	return ErrorResponse(ctx, fiber.StatusInternalServerError, "Internal server error", err.Error())
}
