package dto

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserSignup struct {
	UserLogin
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Phone     string `json:"phone" validate:"required,e164"`
}

type UserVerification struct {
	Code int `json:"code" validate:"required"`
}