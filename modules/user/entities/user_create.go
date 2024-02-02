package userentities

import (
	"math/rand"
	"time"
)

type CreateUser struct {
	FirstName   string    `db:"first_name"`
	LastName    string    `db:"last_name"`
	RoleId      int       `db:"role_id"`
	SsoId       int       `db:"sso_id"`
	UniqueId    string    `db:"unique_id"`
	Email       string    `db:"email"`
	Password    string    `db:"password"`
	PhoneNumber string    `db:"phone_number"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func (cu *CreateUser) AddUniqueId() {
	length := 10
	charset := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	var randomString []rune
	for i := 0; i < length; i++ {
		randomString = append(randomString, charset[rand.Intn(len(charset))])
	}

	cu.UniqueId = string(randomString)
}
