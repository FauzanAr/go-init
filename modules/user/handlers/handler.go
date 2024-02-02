package userhandlers

import (
	"net/http"

	users "github.com/FauzanAr/go-init/modules/user"
	usermodel "github.com/FauzanAr/go-init/modules/user/models"
	"github.com/FauzanAr/go-init/pkg/helper"
	"github.com/FauzanAr/go-init/pkg/logger"
	"github.com/FauzanAr/go-init/pkg/wrapper"
	"github.com/labstack/echo/v4"
)

type UserHandlers struct {
	log logger.Logger
	uc  users.Usecase
}

func NewUserHandlers(log logger.Logger, uc users.Usecase) *UserHandlers {
	return &UserHandlers{
		log: log,
		uc:  uc,
	}
}

func (uh *UserHandlers) Login(c echo.Context) error {
	var req usermodel.LoginUserRequest
	ctx := c.Request().Context()

	if err := c.Bind(&req); err != nil {
		uh.log.Error(ctx, "Error while binding the request", err, nil)
		return wrapper.SendErrorResponse(c, err, nil, http.StatusBadRequest)
	}

	res, err := uh.uc.Login(ctx, req)
	if err != nil {
		return wrapper.SendErrorResponse(c, err, nil, http.StatusBadRequest)
	}

	return wrapper.SendSuccessResponse(c, "Success login", res, http.StatusOK)
}

func (uh *UserHandlers) GetUser(c echo.Context) error {
	var req usermodel.GetUserRequest
	ctx := c.Request().Context()
	user, ok := ctx.Value("user").(*helper.AccessClaims)
	if ok {
		errMsg := wrapper.InternalServerError("Error while converting request")
		return wrapper.SendErrorResponse(c, errMsg, nil, http.StatusInternalServerError)
	}

	req.Id = user.Claims.Id

	res, err := uh.uc.GetUser(ctx, req)
	if err != nil {
		return wrapper.SendErrorResponse(c, err, nil, http.StatusBadRequest)
	}

	return wrapper.SendSuccessResponse(c, "Success", res, http.StatusOK)
}

func (uh *UserHandlers) Register(c echo.Context) error {
	var req usermodel.RegisterUserRequest
	ctx := c.Request().Context()

	if err := c.Bind(&req); err != nil {
		uh.log.Error(ctx, "Error while binding the request", err, nil)
		return wrapper.SendErrorResponse(c, err, nil, http.StatusBadRequest)
	}

	if err := c.Validate(req); err != nil {
		uh.log.Error(ctx, "Error while validate the request", err, nil)
		return wrapper.SendErrorResponse(c, err, nil, http.StatusBadRequest)
	}

	res, err := uh.uc.Register(ctx, req)
	if err != nil {
		return wrapper.SendErrorResponse(c, err, nil, http.StatusBadRequest)
	}

	return wrapper.SendSuccessResponse(c, "Success register", res, http.StatusOK)
}
