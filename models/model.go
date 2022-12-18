package models

import "time"

type User struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Community struct {
	CommunityID   int64     `json:"community_id"`
	CommunityName string    `json:"community_name"`
	Introduction  string    `json:"introduction"`
	CreateTime    time.Time `json:"create_time"`
	UpdateTime    time.Time `json:"update_time"`
}

const (
	StatusExist  = 0
	StatusDelete = 1
)

type Post struct {
	PostID      int64     `json:"post_id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	AuthorID    int64     `json:"author_id"`
	CommunityID int64     `json:"community_id"`
	Status      int       `json:"status"`
	CreateTime  time.Time `json:"create_time"`
	UpdateTime  time.Time `json:"update_time"`
}
