package domain

import "github.com/golang-jwt/jwt/v4"

type JwtCustomClaims struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}
