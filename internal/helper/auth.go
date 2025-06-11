package helper

import (
	"errors"
	"go-ecommerce-app/internal/domain"
	"go-ecommerce-app/pkg/utils"
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
	return Auth{
		Secret: secret,
	}
}

func (a Auth) CreateHashedPassword(password string) (string, error) {
	if len(password) == 0 {
		return "", errors.New("password is required")
	}

	hashP, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// log error
		println(err)
		return "", errors.New("failed to hash password")
	}

	return string(hashP), nil
}

func (a Auth) GenerateToken(id uint, email string, role string) (string, error) {
	if id == 0 || email == "" || role == "" {
		return "", errors.New("id, email, and role are required")
	}

	claims := domain.JwtCustomClaims{
		ID:    id,
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

func (a Auth) VerifyPassword(password string, hashedPassword string) (bool, error) {
	if len(password) == 0 {
		return false, errors.New("password is required")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
			return false, nil
	}
	return true, nil
}

func (a Auth) VerifyToken(bearerToken string) (domain.User, error) {
	tokenString, err := extractTokenString(bearerToken)
	if err != nil {
		return domain.User{}, err
	}

	claims := &domain.JwtCustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(a.Secret), nil
	})
	if err != nil || !token.Valid {
		return domain.User{}, errors.New("invalid or expired token")
	}

	return domain.User{
		ID:       claims.ID,
		Email:    claims.Email,
		UserType: claims.Role,
	}, nil
}

func (a Auth) Authorize(ctx *fiber.Ctx) error {
	user, err := a.VerifyToken(ctx.Get("Authorization"))
	if err != nil {
		println(err)
		return utils.ErrorResponse(ctx, fiber.StatusUnauthorized, "unauthorized", err.Error())
	}

	// store user in context
	ctx.Locals("user", user)

	return ctx.Next()
}

func (a Auth) GetCurrentUser(ctx *fiber.Ctx) domain.User {
	
	claims := ctx.Locals("user").(*domain.JwtCustomClaims)
	return domain.User{
		ID:       claims.ID,
		Email:    claims.Email,
		UserType: claims.Role,
	}
}

func extractTokenString(bearer string) (string, error) {
	parts := strings.Split(bearer, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("invalid token format")
	}
	return parts[1], nil
}

