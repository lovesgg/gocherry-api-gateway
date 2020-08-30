package main

import (
	"fmt"
	"github.com/kataras/iris"
	middleware "gocherry-api-gateway/admin/middware"
	"gocherry-api-gateway/admin/models"
	"gocherry-api-gateway/admin/route"
	"gocherry-api-gateway/admin/services"
	"gocherry-api-gateway/components/common_enum"
)

var appConfig *common_enum.Config

func init() {
	fmt.Println("admin init starting...")

	appConfig = services.GetAppConfig()

	issucc := models.GetInstance().InitDataPool()
	if !issucc {
		fmt.Println("error mysql")
	}

}

func main() {

	app := iris.New()
	app.Use(middleware.NewRecoverPanic())
	app.Use(middleware.LoginAuth)

	route.RegisterRoutes(app)

	_ = app.Run(iris.Addr(":" + appConfig.Admin.AppPort + ""))
}
