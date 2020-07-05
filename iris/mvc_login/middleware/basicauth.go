package middleware

import "github.com/kataras/iris/v12/middleware/basicauth"

//BasicAuth 中间件示例。
var BasicAuth = basicauth.New(
	basicauth.Config{
		Users: map[string]string{
			"admin": "password",
		},
	},
)
