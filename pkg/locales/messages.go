package locales

import "go-ecommerce-app/pkg/errors"


type Language string

const (
	EN Language = "en"
	ID Language = "id"
)

var messages = map[Language]map[errors.ErrorKey]string{
	EN: {
		errors.ErrInputInvalid:    "Invalid input",
		errors.ErrValidationFailed: "Validation failed",
		errors.ErrUserNotFound:    "User not found",
		errors.ErrUserCreationFailed: "Failed to create user",
		errors.ErrUserUpdateFailed: "Failed to update user",
		errors.ErrUserAlreadyVerified: "User already verified",
		errors.ErrUserVerificationInvalid: "User verification invalid",
		errors.ErrUserVerificationExpired: "Code verification expired",
		errors.ErrPasswordInvalid: "Invalid password",
		errors.ErrPasswordHashFailed: "Failed to hash password",
		errors.ErrEmailAlreadyExists: "Email already exists",
	},
	ID: {
		errors.ErrInputInvalid:    "Input tidak valid",
		errors.ErrValidationFailed: "Validasi gagal",
		errors.ErrUserNotFound:    "Pengguna tidak ditemukan",
		errors.ErrUserCreationFailed: "Gagal membuat pengguna",
		errors.ErrUserUpdateFailed: "Gagal memperbarui pengguna",
		errors.ErrUserAlreadyVerified: "Pengguna sudah diverifikasi",
		errors.ErrUserVerificationInvalid: "Verifikasi pengguna tidak valid",
		errors.ErrUserVerificationExpired: "Code verifikasi pengguna telah kadaluarsa",
		errors.ErrPasswordInvalid: "Kata sandi salah",
		errors.ErrPasswordHashFailed: "Gagal mengenkripsi kata sandi",
		errors.ErrEmailAlreadyExists: "Email sudah digunakan",
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
