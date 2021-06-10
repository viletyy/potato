/*
 * @Date: 2021-06-10 18:02:35
 * @LastEditors: viletyy
 * @LastEditTime: 2021-06-10 18:19:12
 * @FilePath: /potato/internal/middleware/translations.go
 */
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/validator/v10"

	ut "github.com/go-playground/universal-translator"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		uni := ut.New(en.New(), zh.New())
		locale := c.GetHeader("locale")
		trans, _ := uni.GetTranslator(locale)
		v, ok := binding.Validator.Engine().(*validator.Validate)
		v.SetTagName("validate")
		if ok {
			switch locale {
			case "zh":
				_ = zhTranslations.RegisterDefaultTranslations(v, trans)
				break
			case "en":
				_ = enTranslations.RegisterDefaultTranslations(v, trans)
				break
			default:
				_ = zhTranslations.RegisterDefaultTranslations(v, trans)
				break
			}
			c.Set("trans", trans)
		}

		c.Next()
	}
}
