package models

import "time"

type User struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Community struct {
	CommunityID   int       `json:"community_id"`
	CommunityName string    `json:"community_name"`
	Introduction  string    `json:"introduction"`
	CreateTime    time.Time `json:"create_time"`
	UpdateTime    time.Time `json:"update_time"`
}
