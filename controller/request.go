package controller

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

const CtxUserIDKey = "userID"

var (
	ErrUserNeedLogin    = errors.New("用户需要登录登录")
	ErrInvalidPageParam = errors.New("页码或页面大小参数错误")
)

// 获取当前用户的ID
func getCurrentUserID(c *gin.Context) (userId int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	// fmt.Println(uid, "--------------")
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

// 获取分页参数page size
func getPageInfo(c *gin.Context) (page, size int64, err error) {
	pageStr := c.Query("page")
	sizeStr := c.Query("size")

	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		err = ErrInvalidPageParam
		page = 1
		return
	}

	size, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		err = ErrInvalidPageParam
		size = 10
		return
	}

	return

}
