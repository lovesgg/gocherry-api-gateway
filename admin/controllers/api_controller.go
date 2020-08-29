package controllers

import "github.com/kataras/iris/context"

var apiController ApiController

type ApiController struct {
	BaseAction
}

func ApiListHandler(ctx context.Context) {
	apiController.GetList(ctx)
}

func ApiSaveHandler(ctx context.Context) {
	apiController.Save(ctx)
}

func ApiGetOneHandler(ctx context.Context) {
	apiController.GetOne(ctx)
}

func ApiDelOneHandler(ctx context.Context) {
	apiController.Del(ctx)
}
