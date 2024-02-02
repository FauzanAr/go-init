package userhandlers

import (
	"github.com/FauzanAr/go-init/pkg/middleware"
	"github.com/labstack/echo/v4"
)

func (uh *UserHandlers) UserRoutes(echo *echo.Group) {
	protectedRoute := echo.Group("")
	protectedRoute.Use(middleware.AuthMiddleware)

	protectedRoute.GET("/v1/user", uh.GetUser)

	echo.POST("/v1/login", uh.Login)
	echo.POST("/v1/register", uh.Register)
}
