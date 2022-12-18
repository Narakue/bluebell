package models

import "time"

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

type ApiPostDetail struct {
	AuthorName string `json:"author_name"`
	*Post      `json:"post"`
	*Community `json:"community"`
}
