package controller

import (
	"encoding/json"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zhaojiasanxing/go_gateway/conf"
	"github.com/zhaojiasanxing/go_gateway/dao"
	"github.com/zhaojiasanxing/go_gateway/dto"
	"github.com/zhaojiasanxing/go_gateway/middleware"
	"github.com/zhaojiasanxing/go_gateway/public"
	"time"
)

type AdminLoginController struct {

}

func AdminLoginRegister(group *gin.RouterGroup)  {
	adminLogin := &AdminLoginController{}
	group.POST("/login", adminLogin.AdminLogin)
	group.GET("/logout", adminLogin.AdminLoginOUt)
}

func (adminLogin *AdminLoginController)AdminLogin(c *gin.Context)  {
	params := &dto.AdminLoginInput{}
	if err := params.BindValidParam(c); err != nil{
		middleware.ResponseErr(c, 1001, err)
		return
	}
	//1、params.UserName 取得管理员信息
	//2、admininfo.salt + params.password sha256 ->saltepassword
	//3、saltpasswrod ==  admininfo.password
	admin := &dao.Admin{}
	admin, err := admin.LoginCheck(conf.DB, params)
	if err != nil{
		middleware.ResponseErr(c, 1002, err)
		return
	}
	//admin.LoginCheck()

	sessInfo := &dto.AdminSessionInfo{
		ID:        admin.Id,
		UserName:  admin.UserName,
		LoginTime: time.Now(),
	}
	//设置session
	sessBts, err := json.Marshal(sessInfo)
	if err != nil{
		middleware.ResponseErr(c, 1003, err)
	}

	sess := sessions.Default(c)
	sess.Set(public.AdminSessionInfoKey, string(sessBts))
	sess.Save()

	out := &dto.AdminLoginOutput{Token:params.UserName}
	middleware.ResponseSuccess(c, out)
}

func (adminLoginOut *AdminLoginController)AdminLoginOUt(c *gin.Context)  {
	sess := sessions.Default(c)
	sess.Delete(public.AdminSessionInfoKey)
	sess.Save()
	middleware.ResponseSuccess(c, "")
}
