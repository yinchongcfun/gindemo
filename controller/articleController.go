package controller

import (
	"gindemo/middleware"
	"gindemo/models"
	"github.com/gin-gonic/gin"
	"log"
)

//文章列
func ArticleList(c *gin.Context) {
	//file, header, e := c.Request.FormFile("upload")
	var article models.Article
	title := c.Request.FormValue("title")
	users, err := article.ListArticle(title)
	if err != nil {
		log.Fatal(err)
	}
	middleware.ResponseSuccess(c, users)
}