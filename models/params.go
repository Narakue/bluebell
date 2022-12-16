package models

type SignUpParam struct {
	Username   string `json:"username" validate:"required"`
	Password   string `json:"password" validate:"required"`
	RePassword string `json:"re_password" validate:"required eqfield=Password"`
}
