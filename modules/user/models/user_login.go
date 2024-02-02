package usermodel

type LoginUserRequest struct {
	Credential string `json:"credential"`
	Password   string `json:"password"`
}
