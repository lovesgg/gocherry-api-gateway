package controllers

import "github.com/kataras/iris/context"

var userController UserController

type UserController struct {
	BaseAction
}

func UserListHandler(ctx context.Context) {
	userController.GetList(ctx)
}


func UserSaveHandler(ctx context.Context) {
	userController.Save(ctx)
}
