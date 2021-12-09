package middleware

import (
	"github.com/gin-gonic/gin"
	"go-project/pkg/api"
)

func Token() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = 200
		token := c.GetHeader("token")
		if token == "" {
			code = 400
			data = "TOKEN NOT FOUND"
		} else {
			// TODO
			if token != "123" {
				code = 404
				data = "TOKEN 验证失败！"
			}
		}

		if code != 200 {
			api.WithCode(code, data).End(c)
			c.Abort()
			return
		}

		c.Next()
	}
}
