package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
)

const CtxUserIDKey = "userID"

var ErrUserNeedLogin = errors.New("用户需要登录登录")

// 获取当前用户的ID
func getCurrentUserID(c *gin.Context) (userId int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrUserNeedLogin
		return
	}
	userId, ok = uid.(int64)
	if !ok {
		err = ErrUserNeedLogin
		return
	}
	return
}
