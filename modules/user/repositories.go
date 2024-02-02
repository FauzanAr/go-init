package users

import (
	"context"

	userentities "github.com/FauzanAr/go-init/modules/user/entities"
)

type Repository interface {
	CreateUser(context.Context, userentities.CreateUser) (int64, error)
	GetUserById(context.Context, int64) ([]userentities.GetUser, error)
	GetUserByEmailOrPhone(context.Context, string) ([]userentities.GetUser, error)
}
