package models

const (
	StatusExist  = 0
	StatusDelete = 1
	OrderTime    = "time"
	OrderScore   = "score"
)

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

type PostParam struct {
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" binding:"required"`
	CommunityID int64  `json:"community_id" binding:"required"`
}

type VoteParam struct {
	PostID int64 `json:"post_id,string" binding:"required"`
	Status *int  `json:"status" binding:"required,oneof=1 0 -1"`
}

type PostListParam struct {
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"page_size" form:"page_size"`
	Order    string `json:"order" form:"order"`
}
