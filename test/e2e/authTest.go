package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Helper: Register user via API
func createTestUser(t *testing.T, email, password string) {
	t.Helper()

	payload := map[string]string{
		"first_name": "Test",
		"last_name":  "User",
		"phone":      "+6285280701948",
		"email":      email,
		"password":   password,
	}

	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/users/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode, "user registration failed")

	printResponseMessage(t, resp)
}

// Helper: Read and print response message
func printResponseMessage(t *testing.T, resp *http.Response) {
	t.Helper()
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	var responseBody map[string]interface{}
	err = json.Unmarshal(bodyBytes, &responseBody)
	assert.NoError(t, err)

	if msg, ok := responseBody["message"]; ok {
		fmt.Printf("Response message: %v\n", msg)
	} else {
		fmt.Println("No message field in response.")
	}
}

func TestRegisterUser(t *testing.T) {
	email := fmt.Sprintf("registeruser_%d@example.com", time.Now().UnixNano())
	password := "12345678"

	t.Run("Register succeeds with valid input", func(t *testing.T) {
		payload := map[string]string{
			"first_name": "Reg",
			"last_name":  "Tester",
			"phone":      "+628111111111",
			"email":      email,
			"password":   password,
		}

		body, _ := json.Marshal(payload)
		req := httptest.NewRequest(http.MethodPost, "/users/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, -1)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.StatusCode, "expected registration to succeed")

		printResponseMessage(t, resp)
	})

	t.Run("Register fails with missing fields", func(t *testing.T) {
		payload := map[string]string{
			"email":    email,
			"password": password,
		}

		body, _ := json.Marshal(payload)
		req := httptest.NewRequest(http.MethodPost, "/users/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, -1)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "expected registration to fail with bad input")

		printResponseMessage(t, resp)
	})

	t.Run("Register fails with duplicate email", func(t *testing.T) {
		payload := map[string]string{
			"first_name": "Reg",
			"last_name":  "Tester",
			"phone":      "+628111111111",
			"email":      email,
			"password":   password,
		}

		body, _ := json.Marshal(payload)
		req := httptest.NewRequest(http.MethodPost, "/users/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, -1)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "expected registration to fail due to duplicate email")

		printResponseMessage(t, resp)
	})
}

func TestLoginUser(t *testing.T) {
	email := fmt.Sprintf("testuser_%d@example.com", time.Now().UnixNano())
	password := "12345678"

	createTestUser(t, email, password)

	t.Run("Login succeeds with valid credentials", func(t *testing.T) {
		payload := map[string]string{
			"email":    email,
			"password": password,
		}

		body, _ := json.Marshal(payload)
		req := httptest.NewRequest(http.MethodPost, "/users/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, -1)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode, "expected login to succeed")

		printResponseMessage(t, resp)
	})

	t.Run("Login fails with incorrect password", func(t *testing.T) {
		payload := map[string]string{
			"email":    email,
			"password": "wrongpassword",
		}

		body, _ := json.Marshal(payload)
		req := httptest.NewRequest(http.MethodPost, "/users/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, -1)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnauthorized, resp.StatusCode, "expected login to fail with incorrect password")

		printResponseMessage(t, resp)
	})
}
