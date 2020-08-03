package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaojiasanxing/go_gateway/dto"
	"github.com/zhaojiasanxing/go_gateway/middleware"
)

type AdminLoginController struct {

}

func AdminLoginRegister(group *gin.RouterGroup)  {
	adminLogin := &AdminLoginController{}
	group.POST("/login", adminLogin.AdminLogin)
}

func (adminLogin *AdminLoginController)AdminLogin(c *gin.Context)  {
	params := &dto.AdminLoginInput{}
	if err := params.BindValidParam(c); err != nil{
		middleware.ResponseErr(c, 1001, err)
		return
	}
	middleware.ResponseSuccess(c, "")
}
