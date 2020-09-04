package main

import (
	"fmt"
	"github.com/kataras/iris"
	middleware "gocherry-api-gateway/admin/middware"
	"gocherry-api-gateway/components/common_enum"
	"gocherry-api-gateway/proxy/filter"
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

	//初始化加载所有server节点
	client, _ := filter.NewClientMon()
	servers, _ := client.GetService()
	fmt.Println(servers)

	route.RegisterRoutes(app, appConfig, servers)

	_ = app.Run(iris.Addr(":"+appConfig.Proxy.AppPort), iris.WithConfiguration(iris.Configuration{
		EnablePathEscape: true,
	}))
}
