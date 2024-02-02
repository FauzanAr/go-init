package users

import (
	"context"

	usermodel "github.com/FauzanAr/go-init/modules/user/models"
)

type Usecase interface {
	Register(context.Context, usermodel.RegisterUserRequest) (usermodel.TokenResponse, error)
	GetUser(context.Context, usermodel.GetUserRequest) (usermodel.GetUserResponse, error)
	Login(context.Context, usermodel.LoginUserRequest) (usermodel.TokenResponse, error)
}
