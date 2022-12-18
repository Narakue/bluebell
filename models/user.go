package models

type User struct {
	UserID   int64  `json:"user_id,string"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
