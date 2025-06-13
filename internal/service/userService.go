package service

import (
	"go-ecommerce-app/internal/domain"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/pkg/locales"
)

//go:generate mockery --name=UserService --dir=internal/service --output=internal/service/mocks --outpkg=mocks
type UserService interface {
	Signup(input dto.UserSignup, lang locales.Language) (string, error)
	Login(input dto.UserLogin, lang locales.Language) (string, error)
	GetVerificationCode(idUser uint, lang locales.Language) (int, error)
	VerifyCode(id uint, code int, lang locales.Language) error
	CreateProfile(id uint, input any) error
	GetProfile(id uint, lang locales.Language) (*domain.User, error)
	UpdateProfile(id uint, input any) error
	BecomeSeller(id uint, input any) (string, error)
	GetCart(id uint) ([]interface{}, error)
	CreateCart(input any, u domain.User) ([]interface{}, error)
	CreateOrder(u domain.User) (int, error)
	GetManyOrder(u domain.User) ([]interface{}, error)
	GetOrderById(id uint, uId uint) ([]interface{}, error)
}
