package controllers

import "github.com/kataras/iris/context"

var appController AppController

type AppController struct {
	BaseAction
}

func AppListHandler(ctx context.Context) {
	appController.GetList(ctx)
}

func AppSaveHandler(ctx context.Context) {
	appController.Save(ctx)
}
