package userusecase

import (
	"context"
	"time"

	users "github.com/FauzanAr/go-init/modules/user"
	userentities "github.com/FauzanAr/go-init/modules/user/entities"
	usermodel "github.com/FauzanAr/go-init/modules/user/models"
	"github.com/FauzanAr/go-init/pkg/helper"
	"github.com/FauzanAr/go-init/pkg/logger"
	"github.com/FauzanAr/go-init/pkg/wrapper"
)

type UserUsecase struct {
	ur  users.Repository
	log logger.Logger
}

func NewUserUsecase(log logger.Logger, ur users.Repository) users.Usecase {
	return UserUsecase{
		ur:  ur,
		log: log,
	}
}

func (u UserUsecase) Register(ctx context.Context, payload usermodel.RegisterUserRequest) (usermodel.TokenResponse, error) {
	var res usermodel.TokenResponse

	hashPassword, err := helper.Hash(payload.Password)
	if err != nil {
		hashError := wrapper.BadRequestError("bad request error")
		u.log.Error(ctx, "Error while hashing password", err, nil)
		return res, hashError
	}

	user := userentities.CreateUser{
		FirstName:   payload.FirstName,
		LastName:    "",
		RoleId:      1,
		SsoId:       0,
		UniqueId:    "",
		Email:       payload.Email,
		Password:    hashPassword,
		PhoneNumber: payload.PhoneNumber,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	user.AddUniqueId()
	id, err := u.ur.CreateUser(ctx, user)
	if err != nil {
		return res, err
	}

	accessToken, err := helper.GenerateAccessToken(ctx, helper.Claims{
		Id:     id,
		Name:   user.FirstName,
		Email:  user.Email,
		Mobile: user.PhoneNumber,
	})

	if err != nil {
		return res, wrapper.BadRequestError("error while generating access token")
	}

	refreshToken, err := helper.GenerateRefreshToken(ctx, helper.Claims{
		Id: id,
	})

	if err != nil {
		return res, wrapper.BadRequestError("error while generating refresh token")
	}

	res.AccessToken = accessToken
	res.RefreshToken = refreshToken

	return res, nil
}

func (u UserUsecase) GetUser(ctx context.Context, payload usermodel.GetUserRequest) (usermodel.GetUserResponse, error) {
	var res usermodel.GetUserResponse
	user, err := u.ur.GetUserById(ctx, payload.Id)
	if err != nil {
		return res, err
	}

	if len(user) > 0 {
		res = usermodel.GetUserResponse{
			FirstName:   user[0].FirstName,
			LastName:    user[0].LastName,
			RoleId:      user[0].RoleId,
			SsoId:       user[0].SsoId,
			UniqueId:    user[0].UniqueId,
			Email:       user[0].Email,
			PhoneNumber: user[0].PhoneNumber,
			CreatedAt:   user[0].CreatedAt,
			UpdatedAt:   user[0].UpdatedAt,
		}
	}

	return res, nil
}

func (u UserUsecase) Login(ctx context.Context, payload usermodel.LoginUserRequest) (usermodel.TokenResponse, error) {
	var res usermodel.TokenResponse
	user, err := u.ur.GetUserByEmailOrPhone(ctx, payload.Credential)
	if err != nil {
		return res, err
	}

	if len(user) == 0 {
		return res, wrapper.NotFoundError("user not valid")
	}

	correctPassword, err := helper.Compare(payload.Password, user[0].Password)
	if err != nil || !correctPassword {
		return res, wrapper.BadRequestError("password incorrect")
	}

	accessToken, err := helper.GenerateAccessToken(ctx, helper.Claims{
		Id:     user[0].Id,
		Name:   user[0].FirstName,
		Email:  user[0].Email,
		Mobile: user[0].PhoneNumber,
	})

	if err != nil {
		return res, wrapper.BadRequestError("error while generating access token")
	}

	refreshToken, err := helper.GenerateRefreshToken(ctx, helper.Claims{
		Id: user[0].Id,
	})

	if err != nil {
		return res, wrapper.BadRequestError("error while generating refresh token")
	}

	res.AccessToken = accessToken
	res.RefreshToken = refreshToken

	return res, nil
}
