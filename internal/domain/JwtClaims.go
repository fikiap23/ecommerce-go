package domain

import "github.com/golang-jwt/jwt/v4"

type JwtCustomClaims struct {
	Sub    uint   `json:"sub"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}
