package jwt

import "github.com/golang-jwt/jwt"

//MyClaims 定义自定义的Claims类
//自定义的用于放userid和username
type MyClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}
