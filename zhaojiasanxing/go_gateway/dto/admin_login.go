package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaojiasanxing/go_gateway/public"
)

type AdminLoginInput struct {
	UserName string `json:"user_name" form:"user_name" comment:"姓名" example:"姓名" validate:"required"`
	Password string `json:"password" form:"password" comment:"密码" example:"密码" validate:"required"`
}

func (param *AdminLoginInput)BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}
