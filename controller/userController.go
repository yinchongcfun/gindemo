package controller

import (
	"errors"
	"fmt"
	"gindemo/dto"
	"gindemo/middleware"
	"gindemo/middleware/jwt"
	"gindemo/models"
	"github.com/gin-gonic/gin"
	"log"
)

func Login(c *gin.Context) {
	loginInput := &dto.LoginInput{}
	if err := loginInput.BindingValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	user := &models.User{}
	fmt.Println(loginInput)
	token, err := user.Login(loginInput.UserName, loginInput.Password)
	if err != nil {
		if err.Error() == "record not found" {
			middleware.ResponseError(c, 500, errors.New("该用户不存在"))
			return
		} else {
			middleware.ResponseError(c, 500, errors.New("登录错误"))
			return
		}
	}
	middleware.ResponseSuccess(c,token)
	return
}

func Register(c *gin.Context) {
	registerInput := &dto.RegisterInput{}
	if err := registerInput.BindingValidParams(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	user := &models.User{}
	user.Name = registerInput.UserName
	user.Password = registerInput.Password
	user.Email = registerInput.Email
	fmt.Printf("email:%s,password:%s,name:%s",user.Email,user.Password,user.Name)
	if err := user.Register(); err != nil {
		middleware.ResponseError(c, 2007, err)
		return
	}
	middleware.ResponseSuccess(c,"注册成功",)
	return
}

func UserList(c *gin.Context) {
	var user models.User
	claims := c.MustGet("claims").(*jwt.CustomClaims)
	users, err := user.ListUsers(claims.Name)
	if err != nil {
		log.Fatal(err)
	}
	middleware.ResponseSuccess(c, users)
}
