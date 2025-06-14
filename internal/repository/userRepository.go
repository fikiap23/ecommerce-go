package repository

import (
	"errors"
	"go-ecommerce-app/internal/domain"
	"log"

	"gorm.io/gorm"
)

//mockery --name=UserRepository --dir=internal/repository --output=internal/service/mocks --outpkg=mocks --case=snake
type UserRepository interface {
	CreateUser(user domain.User) (domain.User, error)
	GetUserByEmail(email string) (domain.User, error)
	GetUserById(id uint) (domain.User, error)
	UpdateUser(id uint, updates *domain.UserUpdatePayload) (*domain.User, error)
}

// userRepository adalah implementasi UserRepository yang menggunakan GORM
type userRepository struct {
	db *gorm.DB
}

// Pastikan di compile-time bahwa *userRepository mengimplementasikan UserRepository
var _ UserRepository = (*userRepository)(nil)

// NewUserRepository membuat instance userRepository baru
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// CreateUser membuat data user baru di database
func (r *userRepository) CreateUser(user domain.User) (domain.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		log.Println("Error creating user:", err)
		return domain.User{}, errors.New("failed to create user")
	}
	return user, err
}

// GetUserByEmail mencari user berdasarkan email
func (r *userRepository) GetUserByEmail(email string) (domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}

// GetUserById mencari user berdasarkan id
func (r *userRepository) GetUserById(id uint) (domain.User, error) {
	var user domain.User
	err := r.db.First(&user, id).Error
	return user, err
}

// UpdateUser mengupdate data user berdasarkan id
func (r *userRepository) UpdateUser(id uint, updates *domain.UserUpdatePayload) (*domain.User, error) {
	user := domain.User{}

	// Cek apakah user ada
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}

	err = r.db.Model(&user).Updates(updates).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
