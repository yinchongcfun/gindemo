package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	//file, _ := os.Create("logs/app/log")
	//gin.DefaultWriter = io.MultiWriter(file)
	router := gin.Default()
	//router.Use(gin.Recovery(), gin.Logger())
	////登录注册
	//router.POST("/login", controller.Login)
	//router.POST("/register", controller.Register)
	////用户相关
	//userRoute := router.Group("user")
	//userRoute.Use(jwt.JWTAuth())
	//userRoute.GET("/list", controller.UserList)
	////文章相关
	//articleRoute := router.Group("article")
	//articleRoute.GET("/list", controller.ArticleList)
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return router
}
