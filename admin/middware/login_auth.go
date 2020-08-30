package middleware

import (
	"github.com/kataras/iris"
	"gocherry-api-gateway/admin/admin_enum"
	"gocherry-api-gateway/admin/controllers"
	"gocherry-api-gateway/components/common_enum"
	"gocherry-api-gateway/components/redis_client"
)

func LoginAuth(ctx iris.Context) {
	url := ctx.Path()
	if url != "/" && url != "/login" {
		auth := ctx.Request().Header.Get("Logintoken")
		loginToken := admin_enum.ACCOUNT_LOGIN_TOKEN + auth
		redis := redis_client.GetProxyRedis()
		value := redis.Get(loginToken)

		if value == "" {
			base := new(controllers.BaseAction)
			base.RenderError(ctx, common_enum.ComError{Msg: "未登录", Code: 1024})
			return
		}
	}

	ctx.Next()
}
