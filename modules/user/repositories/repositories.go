package userrepository

import (
	"context"
	"strings"

	users "github.com/FauzanAr/go-init/modules/user"
	userentities "github.com/FauzanAr/go-init/modules/user/entities"
	mysql "github.com/FauzanAr/go-init/pkg/database"
	"github.com/FauzanAr/go-init/pkg/logger"
	"github.com/FauzanAr/go-init/pkg/wrapper"
)

type UserRepository struct {
	db  *mysql.Mysql
	log logger.Logger
}

func NewUserRepository(log logger.Logger, db *mysql.Mysql) users.Repository {
	return UserRepository{
		db:  db,
		log: log,
	}
}

func (ur UserRepository) CreateUser(ctx context.Context, user userentities.CreateUser) (int64, error) {
	query := `INSERT INTO users (first_name, last_name, role_id, sso_id, unique_id, email, password, phone_number, created_at, updated_at) VALUES (:first_name, :last_name, :role_id, :sso_id, :unique_id, :email, :password, :phone_number, :created_at, :updated_at);`
	id, err := ur.db.CreateByQuery(ctx, query, user)
	if strings.Contains(err.Error(), "Duplicate entry") {
		return 0, wrapper.ValidationError("[Email] or [PhoneNumber] already registered")
	}

	if err != nil {
		return 0, wrapper.BadRequestError("error while saving user")
	}

	return id, nil
}

func (ur UserRepository) GetUserById(ctx context.Context, id int64) ([]userentities.GetUser, error) {
	var res []userentities.GetUser
	query := `SELECT * FROM users WHERE id = ? LIMIT 1`
	err := ur.db.FindByQuery(ctx, query, &res, id)
	if err != nil {
		return []userentities.GetUser{}, wrapper.BadRequestError("error while getting user")
	}

	return res, nil
}

func (ur UserRepository) GetUserByEmailOrPhone(ctx context.Context, emailOrPhone string) ([]userentities.GetUser, error) {
	var res []userentities.GetUser
	query := `SELECT * FROM users WHERE email = ? OR phone_number = ? LIMIT 1`
	err := ur.db.FindByQuery(ctx, query, &res, emailOrPhone, emailOrPhone)
	if err != nil {
		return []userentities.GetUser{}, wrapper.BadRequestError("error while getting user")
	}

	return res, nil
}
