package logic

import (
	"bullbell_test/dao/mysql"
	"bullbell_test/models"
	"bullbell_test/pkg/jwt"
	"bullbell_test/pkg/snowflake"
	"fmt"

	"go.uber.org/zap"
)

// 注册逻辑实现，多个小逻辑叠加的实现
func SignUp(p *models.ParamSignUp) error {
	// 1. 校验y用户是否存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		zap.L().Error(err.Error())

		return err
	}
	// 2. 保存用户信息
	userID := snowflake.GenID()
	user := &models.User{
		UserID:   userID,
		UserName: p.Username,
		Password: p.Password,
	}
	if err := mysql.InsertUser(user); err != nil {
		zap.L().Error(err.Error())
		fmt.Println(err)
		return err
	}
	// 3. 返回错误信息
	return nil
}

// 登录逻辑实现
func Login(p *models.ParamLogin) (*models.User, error) {
	user := &models.User{
		UserName: p.Username,
		Password: p.Password,
	}
	if err := mysql.Login(user); err != nil {

		return nil, err
	}
	// 登录成功，创建token并返回用户信息
	token, err := jwt.GenToken(user.UserID, user.UserName)
	if err != nil {
		return nil, err
	}
	// 返回用户token信息
	user.Token = token
	return user, nil
}
