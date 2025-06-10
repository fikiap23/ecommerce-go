package locales

import "go-ecommerce-app/pkg/errors"


type Language string

const (
	EN Language = "en"
	ID Language = "id"
)

var messages = map[Language]map[errors.ErrorKey]string{
	EN: {
		errors.ErrInvalidInput:    "Invalid input",
		errors.ErrValidationFailed: "Validation failed",
		errors.ErrUserNotFound:    "User not found",
		errors.ErrInvalidPassword: "Invalid password",
		errors.ErrEmailExists:     "Email already exists",
	},
	ID: {
		errors.ErrInvalidInput:    "Input tidak valid",
		errors.ErrValidationFailed: "Validasi gagal",
		errors.ErrUserNotFound:    "Pengguna tidak ditemukan",
		errors.ErrInvalidPassword: "Kata sandi salah",
		errors.ErrEmailExists:     "Email sudah digunakan",
	},
}

func GetMessage(key errors.ErrorKey, lang Language) string {
	if langMap, ok := messages[lang]; ok {
		if msg, exists := langMap[key]; exists {
			return msg
		}
	}
	// fallback to English
	return messages[EN][key]
}
