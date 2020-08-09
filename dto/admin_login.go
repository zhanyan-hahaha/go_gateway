package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaojiasanxing/go_gateway/public"
	"time"
)

type AdminSessionInfo struct {
	ID int `json:"id"`
	UserName string `json:"user_name"`
	LoginTime time.Time `json:"login_time"`
} 

type AdminLoginInput struct {
	UserName string `json:"user_name" form:"user_name" comment:"姓名" example:"姓名" validate:"required,is_valid_username"`
	Password string `json:"password" form:"password" comment:"密码" example:"密码" validate:"required"`
}

func (param *AdminLoginInput)BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

type AdminLoginOutput struct {
	Token string `json:"token" form:"token" comment:"token" validate:""`
}
