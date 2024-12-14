package routes

import (
	"bullbell_test/controller"
	"bullbell_test/logger"
	"bullbell_test/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	//给路由分组
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		v1.POST("/login", controller.LoginHandler)
		v1.POST("/signup", controller.SignUpHandler)
		v1.POST("/community", controller.CommunityHandler)
		v1.POST("/community/:id", controller.CommunityDetailHandler)
		//发帖子
		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		v1.GET("/posts", controller.GetPostListHandler)
		v1.POST("/ping", middlewares.AuthMiddleware(), func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	return r
}
