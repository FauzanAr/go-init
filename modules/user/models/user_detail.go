package usermodel

import "time"

type GetUserRequest struct {
	Id int64
}

type GetUserResponse struct {
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	RoleId      int       `json:"roleId"`
	SsoId       int       `json:"ssoId"`
	UniqueId    string    `json:"uniqueId"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phoneNumber"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
