package usermodel

type RegisterUserRequest struct {
	FirstName   string `json:"firstName" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=6"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
}
