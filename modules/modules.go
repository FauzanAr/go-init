package modules

import (
	"context"

	userhandlers "github.com/FauzanAr/go-init/modules/user/handlers"
	userrepository "github.com/FauzanAr/go-init/modules/user/repositories"
	userusecase "github.com/FauzanAr/go-init/modules/user/usecases"
	mysql "github.com/FauzanAr/go-init/pkg/database"
	"github.com/FauzanAr/go-init/pkg/logger"
	"github.com/labstack/echo/v4"
)

type Modules struct {
	ctx  context.Context
	echo *echo.Echo
	log  logger.Logger
	db   *mysql.Mysql
}

func NewModules(ctx context.Context, echo *echo.Echo, db *mysql.Mysql, log logger.Logger) *Modules {
	return &Modules{
		ctx:  ctx,
		echo: echo,
		log:  log,
		db:   db,
	}
}

func (m *Modules) Init() error {
	m.InitUser()
	return nil
}

func (m *Modules) InitUser() error {
	repository := userrepository.NewUserRepository(m.log, m.db)
	usecase := userusecase.NewUserUsecase(m.log, repository)
	handlers := userhandlers.NewUserHandlers(m.log, usecase)

	group := m.echo.Group("/users")
	handlers.UserRoutes(group)
	return nil
}
