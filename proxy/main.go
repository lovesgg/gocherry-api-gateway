package main

import (
	"fmt"
	"github.com/kataras/iris"
	middleware "gocherry-api-gateway/admin/middware"
	"gocherry-api-gateway/components/common_enum"
	"gocherry-api-gateway/proxy/route"
	services "gocherry-api-gateway/proxy/service"
)

var appConfig *common_enum.Config

func init() {
	fmt.Println("proxy init starting...")
	appConfig = services.GetAppConfig()

}

func main() {
	app := iris.New()
	app.Use(middleware.NewRecoverPanic())

	route.RegisterRoutes(app, appConfig)

	_ = app.Run(iris.Addr(":"+appConfig.Proxy.AppPort), iris.WithConfiguration(iris.Configuration{
		EnablePathEscape: true,
	}))
}
