package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaojiasanxing/go_gateway/controller"
	"github.com/zhaojiasanxing/go_gateway/middleware"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares...)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//demo
	v1 := router.Group("/demo")
	v1.Use(middleware.RequestLog())
	{
		controller.DemoRegister(v1)
	}
	//
	////非登陆接口
	//store := sessions.NewCookieStore([]byte("secret"))
	//apiNormalGroup := router.Group("/api")
	//apiNormalGroup.Use(sessions.Sessions("mysession", store),
	//	middleware.RecoveryMiddleware(),
	//	middleware.RequestLog(),
	//	middleware.TranslationMiddleware())
	//{
	//	controller.ApiRegister(apiNormalGroup)
	//}
	//
	////登陆接口
	//apiAuthGroup := router.Group("/api")
	//apiAuthGroup.Use(
	//	sessions.Sessions("mysession", store),
	//	middleware.RecoveryMiddleware(),
	//	middleware.RequestLog(),
	//	middleware.SessionAuthMiddleware(),
	//	middleware.TranslationMiddleware())
	//{
	//	controller.ApiLoginRegister(apiAuthGroup)
	//}
	return router
}
