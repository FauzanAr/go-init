package userentities

import "time"

type GetUser struct {
	Id          int64     `db:"id"`
	FirstName   string    `db:"first_name"`
	LastName    string    `db:"last_name"`
	RoleId      int       `db:"role_id"`
	SsoId       int       `db:"sso_id"`
	UniqueId    string    `db:"unique_id"`
	Email       string    `db:"email"`
	PhoneNumber string    `db:"phone_number"`
	Password    string    `db:"password"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type Query struct {
	Id    int    `db:"id"`
	Email string `db:"email"`
}
