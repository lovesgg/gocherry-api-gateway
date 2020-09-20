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
	serverNodes := new(filter.ClientMon)
	client, _ := serverNodes.NewClientMon()
	_, _ = client.GetService()
	fmt.Println("alive servers: ",client.ServerList)


	route.RegisterRoutes(app, appConfig, client)

	_ = app.Run(iris.Addr(":"+appConfig.Proxy.AppPort), iris.WithConfiguration(iris.Configuration{
		EnablePathEscape: true,
	}))
}
