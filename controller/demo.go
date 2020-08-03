package controller

import "github.com/gin-gonic/gin"

type DemoController struct {
}

func DemoRegister(router *gin.RouterGroup) {
	demo := DemoController{}
	router.GET("/index", demo.Index)
	//router.Any("/bind", demo.Bind)
	//router.GET("/dao", demo.Dao)
	//router.GET("/redis", demo.Redis)
}


func (demo *DemoController) Index(c *gin.Context) {
	return
}