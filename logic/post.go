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

func GetPostDetailByID(pid int64) (postApiDetail *models.ApiPostDetail, err error) {
	//1.通过帖子post_id,查询帖子
	post := new(models.Post)
	post, err = mysql.GetPostDetailByID(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostDetailByID() failed", zap.Error(err))
		return
	}
	//2.通过post里的auther_id,查询用户
	author, err := mysql.GetAuthorByID(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetAuthorByID() failed", zap.Error(err))
		return
	}
	//3.通过post里的community_id,查询社区
	community, err := mysql.GetCommunityByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityByID() failed", zap.Error(err))
		return
	}
	//4.组装结果
	postApiDetail = &models.ApiPostDetail{
		AuthorName:      author.UserName,
		Post:            post,
		CommunityDetail: community,
	}
	//5.返回结果
	return
}
