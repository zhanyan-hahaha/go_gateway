package middleware

import (
	"github.com/gin-gonic/gin"
	en2 "github.com/go-playground/locales/en"
	zh2 "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/zhaojiasanxing/go_gateway/public"
	"reflect"
)

func TranslationMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		en := en2.New()
		zh := zh2.New()
		uni := ut.New(zh, zh, en)
		val := validator.New()
		locale := context.DefaultQuery("locale", "zh")
		trans, _ := uni.GetTranslator(locale)
		switch locale {
		case "en":
			en_translations.RegisterDefaultTranslations(val, trans)
			val.RegisterTagNameFunc(func(field reflect.StructField) string {
				return field.Tag.Get("en_comment")
			})
			break
		default:
			zh_translations.RegisterDefaultTranslations(val, trans)
			val.RegisterTagNameFunc(func(field reflect.StructField) string {
				return field.Tag.Get("comment")
			})
			val.RegisterValidation("is-validuser", func(fl validator.FieldLevel) bool {
				return fl.Field().String() == "admin"
			})
			val.RegisterTranslation("is-validuser", trans, func(ut ut.Translator) error {
				return ut.Add("is-validuser" , "{0}填写不正确", true)
			}, func(ut ut.Translator, fe validator.FieldError) string {
				t,_ := ut.T("is-validuser", fe.Field())
				return t
			})
			break
		}
		context.Set(public.TranslatorKey, trans)
		context.Set(public.ValidatorKey, val)
		context.Next()
	}
}
