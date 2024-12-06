package controller

import (
	"bullbell_test/logic"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CommunityHandler 处理社区请求
func CommunityHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "community",
	})
	data, err := logic.CommunityList()
	if err != nil {
		zap.L().Error("logic.CommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}

// CommunityDetailHandler 处理社区详情请求
func CommunityDetailHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "community detail",
	})
	originID := c.Param("id")
	id, err := strconv.ParseInt(originID, 10, 64)
	if err != nil {
		zap.L().Error("strconv.Atoi() failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetCommunityDetailByID(id)
	if err != nil {
		zap.L().Error("logic.CommunityDetail() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, data)
}
