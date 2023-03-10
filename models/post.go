package models

import "time"

type Post struct {
	PostID      int64     `json:"post_id,string"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	AuthorID    int64     `json:"author_id,string"`
	CommunityID int64     `json:"community_id,string"`
	Status      int       `json:"status"`
	Score       int       `json:"score"`
	CreateTime  time.Time `json:"create_time"`
	UpdateTime  time.Time `json:"update_time"`
}

type ApiPostDetail struct {
	AuthorName string `json:"author_name"`
	*Post      `json:"post"`
	*Community `json:"community"`
}
