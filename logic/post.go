package logic

import (
	"bullbell_test/dao/mysql"
	"bullbell_test/models"
	"bullbell_test/pkg/snowflake"
	"time"

	"go.uber.org/zap"
)

// CreatePostHandler 创建帖子
func CreatePost(post *models.Post) (err error) {
	//1.添加post_id
	post.ID = int64(snowflake.GenID())
	post.CreateTime = time.Now()
	//2.处理帖子返回结果
	if err = mysql.CreatePost(post); err != nil {
		zap.L().Error("mysql.CreatePost() failed", zap.Error(err))
		return
	}
	return

}
