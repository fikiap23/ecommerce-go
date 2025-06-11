package utils

import (
	"go-ecommerce-app/pkg/errors"
	"go-ecommerce-app/pkg/locales"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Validate = validator.New()

func validationMessage(fe validator.FieldError, lang string) string {
	field := strings.ToLower(fe.Field())
	tag := fe.Tag()

	if msg, ok := ValidationMessages[lang][field][tag]; ok {
		return msg
	}
	if msg, ok := ValidationMessages[lang]["default"][tag]; ok {
		return msg
	}
	if msg, ok := ValidationMessages["en"][field][tag]; ok {
		return msg
	}
	if msg, ok := ValidationMessages["en"]["default"][tag]; ok {
		return msg
	}
	return field + " is invalid"
}


func FormatValidationError(err error, lang string) map[string]string {
	errs := make(map[string]string)
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := strings.ToLower(e.Field())
			errs[field] = validationMessage(e, lang)
		}
	}
	return errs
}

func ParseAndValidate[T any](ctx *fiber.Ctx, dest *T) error {
	lang := GetLangOrDefault(ctx.Get("Accept-Language", "en"))

	if err := ctx.BodyParser(dest); err != nil {
		return NewCustomError(errors.ErrInputInvalid, 400, locales.Language(lang), err.Error())
	}

	if err := Validate.Struct(dest); err != nil {
		validationErrors := FormatValidationError(err, lang)
		return NewCustomError(errors.ErrValidationFailed, 400, locales.Language(lang), validationErrors)
	}

	return nil
}
