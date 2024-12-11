package controller

import (
	"bullbell_test/logic"
	"bullbell_test/models"
	"fmt"

	"github.com/gin-gonic/gin"
	//需要使用V10版本的validator，不然类型断言会报错
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// SignUpHandler 处理注册请求
func SignUpHandler(c *gin.Context) {
	//1.接收参数后验证参数合法性
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		//如果不是validator类型，则直接返回错误信息
		if !ok {
			// c.JSON(http.StatusOK, gin.H{

			// 	"msg": err.Error(),
			// })
			ResponseError(c, CodeInvalidParam)
			return
		}
		//如果是validator类型，则返回错误信息翻译后的信息
		// c.JSON(http.StatusOK, gin.H{
		// 	"msg": removeTopStruct(errs.Translate(trans)),
		// 	// "msg": errs.Translate(trans),
		// })
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	// ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))

	//2.业务处理
	if err := logic.SignUp(p); err != nil {
		// 业务处理失败，返回错误信息
		// c.JSON(http.StatusOK, gin.H{
		// 	"msg": err.Error(),
		// })
		ResponseError(c, CodeServerBusy)
	}
	//4.返回响应
	// c.JSON(http.StatusOK, gin.H{
	// 	"code": 200,
	// })
	ResponseSuccess(c, nil)
}

// LoginHandler 处理登录请求
func LoginHandler(c *gin.Context) {
	//1.接收参数后验证参数合法性
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("Login with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		//如果不是validator类型，则直接返回错误信息
		if !ok {
			// c.JSON(http.StatusOK, gin.H{

			// 	"msg": err.Error(),
			// })
			ResponseError(c, CodeInvalidParam)
			return
		}
		//如果是validator类型，则返回错误信息翻译后的信息
		// c.JSON(http.StatusOK, gin.H{
		// 	"msg": removeTopStruct(errs.Translate(trans)),
		// 	// "msg": errs.Translate(trans),
		// })
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	//2.业务处理
	user, err := logic.Login(p)
	if err != nil {
		// 业务处理失败，返回错误信息
		// c.JSON(http.StatusOK, gin.H{
		// 	"msg": err.Error(),
		// })
		ResponseError(c, CodeServerBusy)
		return
	}
	//4.返回响应
	// c.JSON(http.StatusOK, gin.H{
	// 	"code":    200,
	// 	"user_id": user.UserID,
	// 	"token":   user.Token,
	// })

	//调试代码
	// c.Set(CtxUserIDKey, user.UserID)
	// uid, _ := c.Get(CtxUserIDKey)
	// fmt.Println(uid, "xxxxxxxx")

	ResponseSuccess(c, gin.H{
		"user_id":   fmt.Sprintf("%d", user.UserID), // id值大于1<<53-1  int64类型的最大值是1<<63-1
		"user_name": user.UserName,
		"token":     user.Token,
	})
}
