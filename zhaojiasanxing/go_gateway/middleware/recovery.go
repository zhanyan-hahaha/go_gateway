package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhaojiasanxing/go_gateway/public"
	"runtime/debug"
)

func RecoveryMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil{
				fmt.Println(string(debug.Stack()))
				public.ComLogNotice(context, "_com_panic", map[string]interface{}{
					"error": fmt.Sprint(err),
					"stack": string(debug.Stack()),
				})

			}
		}()
		context.Next()
	}
}
