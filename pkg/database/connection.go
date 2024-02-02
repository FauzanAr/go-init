package mysql

import (
	"context"
	"fmt"
	"time"

	"github.com/FauzanAr/go-init/config"
	"github.com/FauzanAr/go-init/pkg/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Mysql struct {
	ctx context.Context
	cfg config.MySql
	log logger.Logger
	db  *sqlx.DB
}

func NewMysql(ctx context.Context, cfg config.MySql, log logger.Logger) *Mysql {
	return &Mysql{
		ctx: ctx,
		cfg: cfg,
		log: log,
	}
}

func (mysql *Mysql) Connect() (*Mysql, error) {
	host := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", mysql.cfg.Username, mysql.cfg.Password, mysql.cfg.Host, mysql.cfg.Port, mysql.cfg.Name)
	db, err := sqlx.Connect("mysql", host)
	if err != nil {
		mysql.log.Error(mysql.ctx, "Error database", err, nil)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		mysql.log.Error(mysql.ctx, "Error while ping the database", err, nil)
		return nil, err
	}

	db.SetMaxOpenConns(30)
	db.SetMaxIdleConns(15)
	db.SetConnMaxLifetime(10 * time.Minute)

	mysql.db = db

	return mysql, nil
}

func (mysql *Mysql) Close() error {
	return mysql.db.Close()
}
