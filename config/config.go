package config

import (
	"context"
	"fmt"

	"github.com/FauzanAr/go-init/pkg/logger"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	AppEnviroment string `env:"APP_ENVIROMENT,required"`
	AppPort       string `env:"APP_PORT,required"`
	Mysql         MySql
}

type MySql struct {
	Host     string `env:"MYSQL_HOST,required"`
	Port     int    `env:"MYSQL_PORT,required"`
	Password string `env:"MYSQL_PASSWORD,required"`
	Username string `env:"MYSQL_USERNAME,required"`
	Name     string `env:"MYSQL_NAME,required"`
}

func LoadEnv(ctx context.Context, log logger.Logger) (Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Error(ctx, "Error while load enviroment", err, nil)
		return Config{}, err
	}

	var conf Config
	err = env.Parse(&conf)
	if err != nil {
		log.Error(ctx, "Error while parsing the enviroment", err, nil)
		fmt.Println(": ", err)
		return Config{}, err
	}

	return conf, nil
}
