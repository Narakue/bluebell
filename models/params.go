package models

type SignUpParam struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
	Email      string `json:"email" binding:"required"`
}

type LoginParam struct {
	Username string `json:"username"`
	Password string `json:"password" binding:"required"`
}
