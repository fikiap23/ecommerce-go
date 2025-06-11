package helper

import (
	"errors"
	"fmt"
	"go-ecommerce-app/internal/domain"
	"go-ecommerce-app/pkg/utils"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Secret string
}

func SetupAuth(secret string) Auth {
	return Auth{Secret: secret}
}

func (a Auth) CreateHashedPassword(password string) (string, error) {
	if password == "" {
		return "", errors.New("password is required")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Hashing error:", err)
		return "", errors.New("failed to hash password")
	}

	return string(hash), nil
}

func (a Auth) GenerateToken(id uint, email, role string) (string, error) {
	if id == 0 || email == "" || role == "" {
		return "", fmt.Errorf("invalid token payload: id=%d, email=%s, role=%s", id, email, role)
	}

	claims := domain.JwtCustomClaims{
		Sub:    id,
		Email: email,
		Role:  role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(a.Secret))
}

func (a Auth) VerifyPassword(password, hashedPassword string) (bool, error) {
	if password == "" {
		return false, errors.New("password is required")
	}

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, nil // password doesn't match, not an error
	}

	return true, nil
}

func (a Auth) VerifyToken(bearerToken string) (*domain.JwtCustomClaims, error) {
	tokenString, err := extractTokenString(bearerToken)
	if err != nil {
		return nil, err
	}

	claims := &domain.JwtCustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(a.Secret), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}

	return claims, nil
}

func (a Auth) Authorize(ctx *fiber.Ctx) error {
	user, err := a.VerifyToken(ctx.Get("Authorization"))
	if err != nil {
		log.Println("Authorization failed:", err)
		return utils.ErrorResponse(ctx, fiber.StatusUnauthorized, "unauthorized", err.Error())
	}

	ctx.Locals("user", user)
	return ctx.Next()
}

func (a Auth) GetCurrentUser(ctx *fiber.Ctx) *domain.JwtCustomClaims {
	claims, ok := ctx.Locals("user").(*domain.JwtCustomClaims)
	if !ok {
		return nil
	}
	return claims
}

func extractTokenString(bearer string) (string, error) {
	parts := strings.SplitN(bearer, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", errors.New("invalid token format")
	}
	return parts[1], nil
}
