package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaojiasanxing/go_gateway/public"
	"time"
)

type AdminInfoOutput struct {
	ID int `json:"id"`
	UserName string `json:"user_name"`
	LoginTime time.Time `json:"login_time"`
	Avatar string `json:"avatar"`
	Introduction string `json:"introduction"`
	Roles []string `json:"roles"`
}

type ChangePwdInput struct {
	Password string `json:"password" form:"password" validate:"required"`
}

func (param *ChangePwdInput)BIndValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}
