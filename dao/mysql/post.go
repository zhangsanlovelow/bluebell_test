package mysql

import (
	"bullbell_test/models"
	"database/sql"

	"go.uber.org/zap"
)

// CreatePostHandler 创建帖子
func CreatePost(post *models.Post) (err error) {
	sqlStr := `insert into post(
		post_id,community_id,author_id,title,content,create_time) 
		values(?,?,?,?,?,?)
		`
	_, err = db.Exec(sqlStr, post.ID, post.CommunityID, post.AuthorID, post.Title, post.Content, post.CreateTime)

	return
}

// GetPostDetailByID 获取帖子详情
func GetPostDetailByID(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select post_id,community_id,author_id,title,content,create_time from post where post_id = ?`
	post = &models.Post{}
	if err = db.Get(post, sqlStr, pid); err != nil {
		return
	}
	return

}

// GetAuthorByID 通过用户ID获取用户信息
func GetAuthorByID(authorID int64) (author *models.User, err error) {
	author = new(models.User)
	sqlStr := `select user_id,username,email from user where user_id = ?`
	if err = db.Get(author, sqlStr, authorID); err != nil {
		zap.L().Error("db.Get() failed", zap.Error(err))
		if err == sql.ErrNoRows {
			err = nil
			return
		}
		return
	}
	return
}

// GetCommunityByID 通过社区ID获取社区信息
func GetCommunityByID(communityID int64) (CommunityDetail *models.CommunityDetail, err error) {
	CommunityDetail = new(models.CommunityDetail)
	sqlStr := `select community_id,community_name,introduction from community where community_id = ?`
	if err = db.Get(CommunityDetail, sqlStr, communityID); err != nil {
		zap.L().Error("db.Get() failed", zap.Error(err))
		if err == sql.ErrNoRows {
			err = nil
			return
		}
		return
	}
	return
}
