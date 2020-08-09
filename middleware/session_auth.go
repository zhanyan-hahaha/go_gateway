package middleware

import (
	"errors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/zhaojiasanxing/go_gateway/public"
)

func SessionAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		sessions := sessions.Default(context)
		if name, ok := sessions.Get(public.AdminSessionInfoKey).(string); !ok || name == ""{
			ResponseErr(context, InternalErrorCode, errors.New("user not login"))
			context.Abort()
			return
		}
		context.Next()
	}
}

