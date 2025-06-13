package handlers_test

import (
	"bytes"
	"encoding/json"
	"go-ecommerce-app/internal/api/rest/handlers"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/service"
	"go-ecommerce-app/internal/service/mocks"
	"go-ecommerce-app/pkg/errors"
	"go-ecommerce-app/pkg/locales"
	"go-ecommerce-app/pkg/utils"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// Setup app & handler
func setupUserHandler(mockSvc service.UserService) (*fiber.App, *handlers.UserHandler) {
	app := fiber.New()
	h := handlers.NewUserHandler(mockSvc)
	return app, h
}



// --- TEST: Register ---
func TestUserHandler_Register_Success(t *testing.T) {
	mockSvc := new(mocks.UserService)
	app, handler := setupUserHandler(mockSvc)

	input := dto.UserSignup{
		FirstName: "John",
		LastName:  "Doe",
		Phone:     "+6285280701948",
		Email:    "test@example.com",
		Password: "password123",
	}
	token := "mocked-token"

	mockSvc.On("Signup", input, locales.Language("en")).Return(token, nil)


	app.Post("/register", handler.Register)

	body, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Language", "en")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	mockSvc.AssertExpectations(t)
}

func TestUserHandler_Register_Error(t *testing.T) {
	mockSvc := new(mocks.UserService)
	app, handler := setupUserHandler(mockSvc)

	input := dto.UserSignup{
		FirstName: "John",
		LastName:  "Doe",
		Email:    "invalid", // invalid format maybe
		Password: "pass",
		Phone:     "+6285280701948",
	}

	app.Post("/register", handler.Register)

	body, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Language", "en")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

// --- TEST: Login ---
func TestUserHandler_Login_Success(t *testing.T) {
	mockSvc := new(mocks.UserService)
	app, handler := setupUserHandler(mockSvc)

	input := dto.UserLogin{
		Email:    "test@example.com",
		Password: "password123",
	}
	token := "mocked-login-token"

	mockSvc.On("Login", input, locales.Language("en")).Return(token, nil)

	app.Post("/login", handler.Login)

	body, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Language", "en")

	resp, err := app.Test(req)
	assert.NoError(t, err)

	
	// 🔍 Tambahan debug untuk melihat status dan body response
	respBody, _ := io.ReadAll(resp.Body)
	t.Logf("Response Status: %d", resp.StatusCode)
	t.Logf("Response Body: %s", string(respBody))
	
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	mockSvc.AssertExpectations(t)
}

func TestUserHandler_Login_Error(t *testing.T) {
	mockSvc := new(mocks.UserService)
	app, handler := setupUserHandler(mockSvc)

	input := dto.UserLogin{
		Email:    "notfound@example.com",
		Password: "wrongpassword",
	}
	mockSvc.On("Login", input, locales.Language("en")).Return("", utils.NewCustomError(errors.ErrPasswordInvalid, 401, locales.Language("en")))

	app.Post("/login", handler.Login)

	body, _ := json.Marshal(input)
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Language", "en")

	resp, err := app.Test(req)
	assert.NoError(t, err)

	// 🔍 Tambahan debug untuk melihat status dan body response
	respBody, _ := io.ReadAll(resp.Body)
	t.Logf("Response Status: %d", resp.StatusCode)
	t.Logf("Response Body: %s", string(respBody))

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)

	mockSvc.AssertExpectations(t)
}

