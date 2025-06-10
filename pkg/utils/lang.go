package utils

import (
	"go-ecommerce-app/pkg/locales"

	"github.com/gofiber/fiber/v2"
)

// GetLangOrDefault normalizes Accept-Language and returns a supported language code
func GetLangOrDefault(requested string) string {
	supported := map[string]bool{
		"en": true,
		"id": true,
	}

	if len(requested) >= 2 {
		lang := requested[:2]
		if supported[lang] {
			return lang
		}
	}
	return "en"
}

func  GetLanguageFromHeader(ctx *fiber.Ctx) locales.Language {
	langStr := ctx.Get("Accept-Language", string(locales.EN))
	return locales.Language(langStr)
}
