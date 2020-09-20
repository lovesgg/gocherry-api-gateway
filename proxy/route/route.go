package route

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"gocherry-api-gateway/components/common_enum"
	"gocherry-api-gateway/proxy/filter"
)

func RegisterRoutes(app *iris.Application, appConfig *common_enum.Config, servers *filter.ClientMon) {
	app.Any("/*", func(ctx context.Context) {
		proxy := new(filter.ProxyController)
		proxy.DoProxyHandler(ctx, appConfig, servers)
	})
}
