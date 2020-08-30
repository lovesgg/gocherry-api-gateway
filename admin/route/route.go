package route

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/core/router"
	"gocherry-api-gateway/admin/controllers"
)

func RegisterRoutes(app *iris.Application) {
	app.Get("/", func(ctx context.Context) {
		_, _ = ctx.WriteString("admin 200")
	})

	app.Post("/login", controllers.LoginHandler)

	/**
	应用级
	*/
	app.PartyFunc("/app", func(route router.Party) {
		route.Any("/list", controllers.AppListHandler)
		route.Post("/save", controllers.AppSaveHandler)
	})

	/**
	集群（服务）
	*/
	app.PartyFunc("/cluster", func(route router.Party) {
		route.Post("/list", controllers.ClusterListHandler)
		route.Post("/save", controllers.ClusterSaveHandler)
	})

	/**
	服务的分布节点
	*/
	app.PartyFunc("/server", func(route router.Party) {
		route.Post("/list", controllers.ServerListHandler)
		route.Post("/save", controllers.ServerSaveHandler)
		route.Post("/del", controllers.ServerDelHandler)
	})

	/**
	api接口
	*/
	app.PartyFunc("/api", func(route router.Party) {
		route.Post("/list", controllers.ApiListHandler)
		route.Post("/save", controllers.ApiSaveHandler)
		route.Post("/get_one", controllers.ApiGetOneHandler)
		route.Post("/del", controllers.ApiDelOneHandler)
	})

	/**
	用户接口
	*/
	app.PartyFunc("/user", func(route router.Party) {
		route.Post("/list", controllers.UserListHandler)
		route.Post("/save", controllers.UserSaveHandler)
		route.Post("/del", controllers.UserDelHandler)
	})

	app.PartyFunc("/admin", func(route router.Party) {
		//加管理员中间件 只有管理员才可操作 比如用户的添加删除 服务的删除等

		route.Post("/user_add", func(ctx context.Context) {

		})
	})

}
