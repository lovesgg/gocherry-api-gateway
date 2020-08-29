package controllers

import "github.com/kataras/iris/context"

var serverController ServerController

type ServerController struct {
	BaseAction
}

func ServerListHandler(ctx context.Context) {
	serverController.GetList(ctx)
}

func ServerSaveHandler(ctx context.Context) {
	serverController.Save(ctx)
}

func ServerDelHandler(ctx context.Context) {
	serverController.Del(ctx)
}
