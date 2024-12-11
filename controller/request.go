package controller

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

const CtxUserIDKey = "userID"

var ErrUserNeedLogin = errors.New("用户需要登录登录")

// 获取当前用户的ID
func getCurrentUserID(c *gin.Context) (userId int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	fmt.Println(uid, "--------------")
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
