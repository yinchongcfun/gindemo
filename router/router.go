package router

import (
	"gindemo/controller"
	"gindemo/middleware/jwt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func InitRouter() *gin.Engine {
	file, _ := os.Create("logs/app/log")
	gin.DefaultWriter = io.MultiWriter(file)
	router := gin.Default()
	router.Use(gin.Recovery(), gin.Logger())
	//登录注册
	router.POST("/login", controller.Login)
	router.POST("/register", controller.Register)
	//用户相关
	userRoute := router.Group("user")
	userRoute.Use(jwt.JWTAuth())
	userRoute.GET("/list", controller.UserList)
	//文章相关
	articleRoute := router.Group("article")
	articleRoute.GET("/list", controller.ArticleList)
	return router
}
