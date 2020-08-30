package middleware

import (
	"encoding/json"
	"github.com/kataras/iris"
	"gocherry-api-gateway/admin/admin_enum"
	"gocherry-api-gateway/admin/models"
	"gocherry-api-gateway/components/log_client"
	"gocherry-api-gateway/components/redis_client"
	"gocherry-api-gateway/components/utils"
)

/**
后台的任何操作记录 任何路由
*/

func OperateLog(ctx iris.Context) {
	var account models.AdminAccount
	url := ctx.Path()
	if url != "/" && url != "/login" {
		auth := ctx.Request().Header.Get("Logintoken")
		loginToken := admin_enum.ACCOUNT_LOGIN_TOKEN + auth
		redis := redis_client.GetProxyRedis()
		value := redis.Get(loginToken)

		_ = json.Unmarshal([]byte(value), &account)

		log_client.LogInfo(ctx, map[string]interface{}{
			"operate_id": account.Phone,
			"url":        url,
			"time":       utils.GetNowTimeFormat(),
		})

	}

	ctx.Next()
}
