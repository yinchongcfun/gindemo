package models

import (
	"errors"
	"gindemo/database"
	"gindemo/middleware/jwt"
	jwtgo "github.com/dgrijalva/jwt-go" //需要安装 然后调用这个jwt-go包
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type User struct {
	Id       int    `form:"id" json:"id" gorm:"PRIMARY_KEY"`
	Name     string `form:"username" json:"username"`
	Email    string `form:"email" json:"email",binding:"required"`
	Password string `form:"password" json:"-",binding:"required"`
}

type LoginResult struct {
	User  interface{} `json:"user"`
	Token string `json:"token"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) Login(name string, password string) (token LoginResult, err error) {
	var user User
	obj := database.GormPool.Where("name = ? and password=?", name, password).First(&user)
	if err = obj.Error; err != nil {
		return
	}
	generateToken := GenerateToken(user)
	return generateToken, nil
}

func (u *User) Register() error {
	//先判断email是否为已经存在
	find := database.GormPool.Where("email=?", u.Email)
	if find != nil {
		return errors.New("email已经存在了")
	}
	if err := database.GormPool.Save(u).Error; err != nil {
		return errors.New("注册失败")
	}
	return nil
}

func (u *User) ListUsers(name string) (users []User, err error) {
	query := database.GormPool
	if name != "" {
		query = query.Where("name=?", name)
	}
	err = query.Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return
}

// 生成令牌  创建jwt风格的token
func GenerateToken(user User) LoginResult {
	j := &jwt.JWT{
		[]byte("newtrekWang"),
	}
	claims := jwt.CustomClaims{
		user.Id,
		user.Name,
		user.Password,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "newtrekWang",                   //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)
	if err != nil {
		return LoginResult{
			User:  user,
			Token: token,
		}
	}

	log.Println(token)
	data := LoginResult{
		User:  user,
		Token: token,
	}
	return data
}
