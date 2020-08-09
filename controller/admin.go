package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zhaojiasanxing/go_gateway/conf"
	"github.com/zhaojiasanxing/go_gateway/dao"
	"github.com/zhaojiasanxing/go_gateway/dto"
	"github.com/zhaojiasanxing/go_gateway/middleware"
	"github.com/zhaojiasanxing/go_gateway/public"
)

type AdminController struct {

}

func AdminRegister(group *gin.RouterGroup)  {
	adminController := &AdminController{}
	group.GET("/admin_info",adminController.AdminInfo)
	group.POST("/change_pwd", adminController.ChangePwd)
}

func (admin *AdminController)AdminInfo(c *gin.Context)  {

	sess := sessions.Default(c)
	sessionInfo := sess.Get(public.AdminSessionInfoKey)
	sessInfoStr := sessionInfo.(string)
	adminSessionInfo := &dto.AdminSessionInfo{}
	err := json.Unmarshal([]byte(sessInfoStr), adminSessionInfo)
	if err != nil{
		middleware.ResponseErr(c, 2000, err)
		return
	}

	out := &dto.AdminInfoOutput{
		ID:           adminSessionInfo.ID,
		UserName:     adminSessionInfo.UserName,
		LoginTime:    adminSessionInfo.LoginTime,
		Avatar:       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		Introduction: "i am a super administrator",
		Roles:        []string{"admin"},
	}
	middleware.ResponseSuccess(c, out)
}

func (adminLogin *AdminController)ChangePwd(c *gin.Context)  {
	param := &dto.ChangePwdInput{}
	if err := param.BIndValidParam(c); err != nil{
		middleware.ResponseErr(c, 3000, err)
		return
	}
	//1、session读取用户信息到结构体 sessInfo
	//2、sessInfo.ID读取数据库，adminInfo
	//3、params.password + adminInfo.salt sha256 saltPassword
	//4、saltPassword save adminInfo.save

	sess := sessions.Default(c)
	sessInfo := sess.Get(public.AdminSessionInfoKey)
	sessInfoStr := sessInfo.(string)
	adminSessionInfo := &dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(fmt.Sprintf(sessInfoStr)), adminSessionInfo); err != nil{
		middleware.ResponseErr(c, 3001, err)
		return
	}
	adminInfo := &dao.Admin{}
	adminInfo, err := adminInfo.Find(conf.DB, (&dao.Admin{UserName:adminSessionInfo.UserName}))
	if err != nil{
		middleware.ResponseErr(c, 3002, err)
	}

	saltPassword := public.GinSaltPassword(adminInfo.Salt, param.Password)
	adminInfo.Password = saltPassword
	err = adminInfo.Save(conf.DB)
	if err != nil{
		middleware.ResponseErr(c, 3003, err)
		return
	}
	middleware.ResponseSuccess(c, "")
}
