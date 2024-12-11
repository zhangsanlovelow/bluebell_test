package models

import "time"

//注意内存对齐

// Post是帖子模型，用户提交帖子标题和内容
type Post struct {
	ID          int64     `json:"post_id" db:"post_id"`                              //帖子ID
	AuthorID    int64     `json:"author_id" db:"author_id"`                          //作者ID
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"` //社区ID
	Status      int64     `json:"status" db:"status"`                                //帖子状态
	Title       string    `json:"title" db:"title" binding:"required"`               //帖子标题
	Content     string    `json:"content" db:"content" binding:"required"`           //帖子内容
	CreateTime  time.Time `json:"create_time" db:"create_time"`                      //创建时间
	UpdateTime  time.Time `json:"update_time" db:"update_time"`                      //	更新时间
}

// ApiPostDetail 是帖子详情模型，包含帖子信息和社区信息
type ApiPostDetail struct {
	AuthorName string `json:"author_name"`
	*Post
	*CommunityDetail `json:"community"`
}
