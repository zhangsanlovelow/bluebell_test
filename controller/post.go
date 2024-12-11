package controller

import (
	"bullbell_test/logic"
	"bullbell_test/models"
	"strconv"

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

// GetPostDetailHandler 通过帖子id获取帖子详情
func GetPostDetailHandler(c *gin.Context) {

	//1.获取帖子id
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("strconv.Atoi() failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//2.通过帖子id获取帖子详情
	data, err := logic.GetPostDetailByID(pid)
	if err != nil {
		zap.L().Error("logic.GetPostDetailByID() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	//3.返回
	ResponseSuccess(c, data)
}
