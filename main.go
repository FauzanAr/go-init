package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/FauzanAr/go-init/config"
	"github.com/FauzanAr/go-init/modules"
	mysql "github.com/FauzanAr/go-init/pkg/database"
	"github.com/FauzanAr/go-init/pkg/logger"
	"github.com/FauzanAr/go-init/pkg/middleware"
	"github.com/FauzanAr/go-init/pkg/validator"
	"github.com/FauzanAr/go-init/pkg/wrapper"

	"github.com/labstack/echo/v4"
)

func main() {
	ctx := context.Background()
	log := logger.NewLogger()
	conf, err := config.LoadEnv(ctx, log)
	if err != nil {
		panic("Error while loading enviroment")
	}

	server := echo.New()
	server.Use(middleware.EchoRequestTrace(log))
	server.Validator = validator.NewValidator()

	mysqlDb := mysql.NewMysql(ctx, conf.Mysql, log)
	mysql, err := mysqlDb.Connect()
	if err != nil {
		panic("Database error")
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM)
	signal.Notify(quit, syscall.SIGINT)

	go func() {
		<-quit
		log.Error(ctx, "Server is shutting down...", nil, nil)

		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		mysql.Close()
		log.Sync()
		server.Shutdown(ctx)
	}()

	server.GET("/", func(c echo.Context) error {
		return wrapper.SendSuccessResponse(c, "Success", map[string]string{"message": "Server running"}, 200)
	})

	modules.NewModules(ctx, server, mysql, log).Init()

	server.Logger.Fatal(server.Start(":" + conf.AppPort))
}
