package service

import (
	"go-ecommerce-app/internal/domain"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/repository"
	"go-ecommerce-app/pkg/errors"
	"go-ecommerce-app/pkg/locales"
	"go-ecommerce-app/pkg/utils"
	"log"
)

type UserService struct {
	Repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
    return &UserService{Repo: repo}
}


func (s *UserService) Signup(input dto.UserSignup, lang locales.Language) (string, error) {
	log.Println("Signup input:", input)

	// Cek apakah email sudah terdaftar
	existingUser, err := s.Repo.GetUserByEmail(input.Email)
	if err == nil && existingUser.ID != 0 {
		return "", utils.NewCustomError(errors.ErrEmailExists, 400, lang)
	}

	// Buat user baru domain.User dari input dto.UserSignup
	newUser := domain.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email: input.Email,
		Phone: input.Phone,
		Password: input.Password, // NOTE: seharusnya hash dulu password nya
	}

	createdUser, err := s.Repo.CreateUser(newUser)
	if err != nil {
		return "", err
	}

	// Generate token dummy (bisa diganti JWT atau lainnya)
	token := "dummy_token_for_user_" + createdUser.Email

	return token, nil
}

func (s *UserService) Login(input dto.UserLogin, lang locales.Language) (string, error) {
	user, err := s.Repo.GetUserByEmail(input.Email)
	if err != nil {
		return "", utils.NewCustomError(errors.ErrUserNotFound, 404, lang)
	}

	if user.Password != input.Password {
		return "", utils.NewCustomError(errors.ErrInvalidPassword, 401, lang)
	}

	token := "dummy_token_for_user_" + user.Email
	return token, nil
}

func (s *UserService) GetVerificationCode(e domain.User) (int, error) {
	// logic
	return 0, nil
}

func (s *UserService) VerifyCode(id uint, code int) error {
	// logic
	return nil
}

func (s *UserService) CreateProfile(id uint, input any) error {
	// logic
	return nil
}

func (s *UserService) GetProfile(id uint) (*domain.User, error) {
	// logic
	return nil, nil
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