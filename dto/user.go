package dto

import (
	"errors"
	"gindemo/public"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
	"strings"
)

type LoginInput struct {
	UserName string `form:"username" validate:"required"`
	Password string `form:"password" validate:"required"`
}


func (o *LoginInput) BindingValidParams(c *gin.Context) error {
	//绑定数据
	if err := c.ShouldBind(o); err != nil {
		return err
	}
	v := c.Value("trans")
	trans, ok := v.(ut.Translator)
	if !ok {
		trans, _ = public.Uni.GetTranslator("zh")
	}
	////验证器注册翻译器
	//e := zh_translations.RegisterDefaultTranslations(public.Validate, trans)
	//if e != nil {
	//	return e
	//}
	//验证
	err := public.Validate.Struct(o)
	if err != nil {
		sliceErrs := []string{}
		for _, e := range err.(validator.ValidationErrors) {

			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return errors.New(strings.Join(sliceErrs, ","))
	}
	return nil
}

type RegisterInput struct {
	UserName string `form:"username"`
	Password string `form:"password" validate:"required"`
	Email    string `form:"email" validate:"required"`
}

func (o *RegisterInput) BindingValidParams(c *gin.Context) error {
	if err := c.ShouldBind(o); err != nil {
		return err
	}
	v := c.Value("trans")
	trans, ok := v.(ut.Translator)
	if !ok {
		trans, _ = public.Uni.GetTranslator("zh")
	}
	//验证器注册翻译器
	zh_translations.RegisterDefaultTranslations(public.Validate, trans)
	err := public.Validate.Struct(o)
	if err != nil {
		sliceErrs := []string{}
		for _, e := range err.(validator.ValidationErrors) {

			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return errors.New(strings.Join(sliceErrs, ","))
	}
	return nil
}


