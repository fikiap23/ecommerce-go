package service_test

import (
	"errors"
	"go-ecommerce-app/internal/domain"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/service"
	"go-ecommerce-app/internal/service/mocks"
	"go-ecommerce-app/pkg/locales"
	"go-ecommerce-app/pkg/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserService_Login_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	mockAuth := new(mocks.Auth)
	svc := service.NewUserService(mockRepo, mockAuth)

	input := dto.UserLogin{
		Email:    "test@example.com",
		Password: "password123",
	}

	user := domain.User{
		ID:       1,
		Email:    input.Email,
		Password: "hashed-password",
	}

	mockRepo.On("GetUserByEmail", input.Email).Return(user, nil)
	mockAuth.On("VerifyPassword", input.Password, user.Password).Return(true, nil)
	mockAuth.On("GenerateToken", user.ID, user.Email, user.UserType).Return("mocked-token", nil)

	token, err := svc.Login(input, locales.EN)

	assert.NoError(t, err)
	assert.Equal(t, "mocked-token", token)
	mockRepo.AssertExpectations(t)
	mockAuth.AssertExpectations(t)
}

func TestUserService_Login_InvalidPassword(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	mockAuth := new(mocks.Auth)
	svc := service.NewUserService(mockRepo, mockAuth)

	input := dto.UserLogin{
		Email:    "test@example.com",
		Password: "wrongpass",
	}

	user := domain.User{
		ID:       1,
		Email:    input.Email,
		Password: "hashed-password",
	}

	mockRepo.On("GetUserByEmail", input.Email).Return(user, nil)
	mockAuth.On("VerifyPassword", input.Password, user.Password).Return(false, nil)

	token, err := svc.Login(input, locales.EN)

	assert.Error(t, err)
	assert.Empty(t, token)

	customErr, ok := err.(*utils.CustomError)
	assert.True(t, ok)
	assert.Equal(t, 401, customErr.StatusCode)
	mockRepo.AssertExpectations(t)
	mockAuth.AssertExpectations(t)
}

func TestUserService_Login_UserNotFound(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	mockAuth := new(mocks.Auth)
	svc := service.NewUserService(mockRepo, mockAuth)

	input := dto.UserLogin{
		Email:    "notfound@example.com",
		Password: "whatever",
	}

	mockRepo.On("GetUserByEmail", input.Email).Return(domain.User{}, errors.New("not found"))

	token, err := svc.Login(input, locales.EN)

	assert.Error(t, err)
	assert.Empty(t, token)

	customErr, ok := err.(*utils.CustomError)
	assert.True(t, ok)
	assert.Equal(t, 404, customErr.StatusCode)
	mockRepo.AssertExpectations(t)
}

func TestUserService_Signup_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	mockAuth := new(mocks.Auth)
	svc := service.NewUserService(mockRepo, mockAuth)

	input := dto.UserSignup{
		FirstName: "John",
		LastName:  "Doe",
		Phone:     "08123456789",
		Email:     "john@example.com",
		Password:  "plain-password",
	}

	mockRepo.On("GetUserByEmail", input.Email).Return(domain.User{}, nil)

	hashed := "hashed-password"
	mockAuth.On("CreateHashedPassword", input.Password).Return(hashed, nil)

	newUser:= domain.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Phone:     input.Phone,
		Password:  hashed,
	}

	mockRepo.On("CreateUser", newUser).Return(domain.User{
		ID:        1, 
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Email:     newUser.Email,
		Phone:     newUser.Phone,
		Password:  newUser.Password,
		UserType:  "customer",
	}, nil)

	
	mockAuth.On("GenerateToken", uint(1), input.Email, "customer").Return("signup-token", nil)

	token, err := svc.Signup(input, locales.EN)

	assert.NoError(t, err)
	assert.Equal(t, "signup-token", token)
	mockRepo.AssertExpectations(t)
	mockAuth.AssertExpectations(t)
}
