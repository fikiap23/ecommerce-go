package service

import (
	"go-ecommerce-app/internal/domain"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/helper"
	"go-ecommerce-app/internal/repository"
	"go-ecommerce-app/pkg/errors"
	"go-ecommerce-app/pkg/locales"
	"go-ecommerce-app/pkg/utils"
	"time"
)

type UserService struct {
	Repo repository.UserRepository
	Auth helper.Auth
}

func NewUserService(repo repository.UserRepository, auth helper.Auth) *UserService {
    return &UserService{Repo: repo, Auth: auth}
}


func (s *UserService) Signup(input dto.UserSignup, lang locales.Language) (string, error) {
	// Cek apakah email sudah terdaftar
	existingUser, err := s.Repo.GetUserByEmail(input.Email)
	if err == nil && existingUser.ID != 0 {
		return "", utils.NewCustomError(errors.ErrEmailAlreadyExists, 400, lang)
	}

	hashedPassword, err := s.Auth.CreateHashedPassword(input.Password)
	if err != nil {
		return "", utils.NewCustomError(errors.ErrPasswordHashFailed, 400, lang)
	}

	// Buat user baru domain.User dari input dto.UserSignup
	newUser := domain.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email: input.Email,
		Phone: input.Phone,
		Password: hashedPassword,
	}

	createdUser, err := s.Repo.CreateUser(newUser)
	if err != nil {
		return "", utils.NewCustomError(errors.ErrUserCreationFailed, 500, lang, err.Error())
	}

	// Generate token
	token, err := s.Auth.GenerateToken(createdUser.ID, createdUser.Email, createdUser.UserType)
	if err != nil {
		return "", utils.NewCustomError(errors.ErrUserCreationFailed, 500, lang, err.Error())
	}

	return token, nil
}

func (s *UserService) Login(input dto.UserLogin, lang locales.Language) (string, error) {
	user, err := s.Repo.GetUserByEmail(input.Email)
	if err != nil {
		return "", utils.NewCustomError(errors.ErrUserNotFound, 404, lang)
	}

	// verify password
	match, err := s.Auth.VerifyPassword(input.Password, user.Password)
	if err != nil || !match {
		return "", utils.NewCustomError(errors.ErrPasswordInvalid, 401, lang)
	}

	// generate token
	token, err := s.Auth.GenerateToken(user.ID, user.Email, user.UserType)
	if err != nil {
		return "", utils.NewCustomError(errors.ErrPasswordInvalid, 401, lang)
	}

	return token, nil
}

func (s *UserService) isVerified(id uint) bool {
	currentUser,err:= s.Repo.GetUserById(id)

	return err == nil && currentUser.Verified
}

func (s *UserService) GetVerificationCode(idUser uint) (int, error) {
	// if user already verified
	if s.isVerified(idUser) {
		return 0, nil
	}

	// generate code
	code := s.Auth.GenerateCode()	

	// update user
	expiresAt := time.Now().Add(5 * time.Minute)

	_, err := s.Repo.UpdateUser(idUser, domain.UserUpdatePayload{
		Code:      &code,
		ExpiresAt: &expiresAt,
	})
	if err != nil {
		return 0, err
	}

	// return verification code
	return code, nil
}

func (s *UserService) VerifyCode(id uint, code int) error {
	// logic
	return nil
}

func (s *UserService) CreateProfile(id uint, input any) error {
	// logic
	return nil
}

func (s *UserService) GetProfile(id uint, lang locales.Language) (*domain.User, error) {
	user, err := s.Repo.GetUserById(id)

	if err != nil {
		return nil, utils.NewCustomError(errors.ErrUserNotFound, 404, lang)
	}
	return &user, nil
}

func (s *UserService) UpdateProfile(id uint, input any) error {
	// logic
	return nil
}

func (s *UserService) BecomeSeller(id uint, input any) (string, error) {
	// logic
	return "", nil
}

func (s *UserService) GetCart(id uint) ([]interface{}, error) {
	// logic
	return nil, nil
}

func (s *UserService) CreateCart(input any, u domain.User) ([]interface{}, error) {
	// logic
	return nil, nil
}

func (s *UserService) CreateOrder( u domain.User) (int, error) {
	// logic
	return 0, nil
}

func (s *UserService) GetManyOrder(u domain.User) ([]interface{}, error) {
	// logic
	return nil, nil
}

func (s *UserService) GetOrderById(id uint, uId uint) ([]interface{}, error) {
	// logic
	return nil, nil
}