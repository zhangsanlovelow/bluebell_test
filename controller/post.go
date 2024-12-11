package controller

import (
	"bullbell_test/logic"
	"bullbell_test/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreatePostHandler 创建帖子
func CreatePostHandler(c *gin.Context) {
	//1.获取参数并校验
	post := new(models.Post)
	if err := c.ShouldBind(post); err != nil {
		zap.L().Error("c.ShouldBind() failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//2.取出当前用户id
	// userId, err := getCurrentUserID(c)
	// if err != nil {
	// 	zap.L().Error("getCurrentUserID() failed", zap.Error(err))
	// 	ResponseError(c, CodeNeedLogin)
	// 	return
	// }
	// post.AuthorID = userId
	//3.处理帖子数据
	if err := logic.CreatePost(post); err != nil {
		zap.L().Error("logic.CreatePost() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//4.返回结果
	ResponseSuccess(c, nil)
}
