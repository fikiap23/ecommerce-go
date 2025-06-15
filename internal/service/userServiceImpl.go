package service

import (
	"fmt"
	"go-ecommerce-app/config"
	"go-ecommerce-app/internal/domain"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/helper"
	"go-ecommerce-app/internal/repository"
	"go-ecommerce-app/pkg/errors"
	"go-ecommerce-app/pkg/locales"
	"go-ecommerce-app/pkg/notification"
	"go-ecommerce-app/pkg/utils"
	"time"
)

type userService struct {
	Repo repository.UserRepository
	Auth helper.Auth
	Config config.AppConfig
}

func NewUserService(repo repository.UserRepository, auth helper.Auth, config config.AppConfig) UserService {
	return &userService{Repo: repo, Auth: auth, Config: config}
}

func (s *userService) Signup(input dto.UserSignup, lang locales.Language) (string, error) {
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

func (s *userService) Login(input dto.UserLogin, lang locales.Language) (string, error) {
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

func (s *userService) isVerified(id uint) bool {
	currentUser,err:= s.Repo.GetUserById(id)

	return err == nil && currentUser.Verified
}

func (s *userService) GetVerificationCode(idUser uint, lang locales.Language)  error {
	// if user already verified
	if s.isVerified(idUser) {
		return utils.NewCustomError(errors.ErrUserAlreadyVerified, 400, lang)
	}

	// generate code
	code := s.Auth.GenerateCode()	

	// update user
	user, err := s.Repo.UpdateUser(idUser, &domain.UserUpdatePayload{
		Code:      &code,
		ExpiresAt: utils.PtrTime(time.Now().Add(5 * time.Minute)),
	})

	if err != nil {
		return utils.NewCustomError(errors.ErrUserUpdateFailed, 500, lang, err.Error())
	}

	// send email
	notificationClient := notification.NewNotificationClient(s.Config)
	message := fmt.Sprintf("Hi %s, your verification code is %d", user.FirstName, code)
	notificationClient.SendSMS(user.Phone, message)

	// return verification code
	return nil
}

func (s *userService) VerifyCode(id uint, code int, lang locales.Language) error {
	// get user
	user, err := s.Repo.GetUserById(id)
	if err != nil {
		return utils.NewCustomError(errors.ErrUserNotFound, 404, lang)
	}

	if(user.Verified) {
		return utils.NewCustomError(errors.ErrUserAlreadyVerified, 400, lang)
	}

	if(user.Code != code) {
		return utils.NewCustomError(errors.ErrUserVerificationInvalid, 400, lang)
	}

	if(user.ExpiresAt.Before(time.Now())) {
		return utils.NewCustomError(errors.ErrUserVerificationExpired, 400, lang)
	}

	// update user
	_, err = s.Repo.UpdateUser(id, &domain.UserUpdatePayload{
		Verified: utils.PtrBool(true),
		Code:     utils.PtrInt(0),
	})

	if err != nil {
		return utils.NewCustomError(errors.ErrUserUpdateFailed, 500, lang, err.Error())
	}
	
	return nil
}

func (s *userService) CreateProfile(id uint, input any) error {
	// logic
	return nil
}

func (s *userService) GetProfile(id uint, lang locales.Language) (*domain.User, error) {
	user, err := s.Repo.GetUserById(id)

	if err != nil {
		return nil, utils.NewCustomError(errors.ErrUserNotFound, 404, lang)
	}
	return &user, nil
}

func (s *userService) UpdateProfile(id uint, input any) error {
	// logic
	return nil
}

func (s *userService) BecomeSeller(id uint, input any) (string, error) {
	// logic
	return "", nil
}

func (s *userService) GetCart(id uint) ([]interface{}, error) {
	// logic
	return nil, nil
}

func (s *userService) CreateCart(input any, u domain.User) ([]interface{}, error) {
	// logic
	return nil, nil
}

func (s *userService) CreateOrder( u domain.User) (int, error) {
	// logic
	return 0, nil
}

func (s *userService) GetManyOrder(u domain.User) ([]interface{}, error) {
	// logic
	return nil, nil
}

func (s *userService) GetOrderById(id uint, uId uint) ([]interface{}, error) {
	// logic
	return nil, nil
}