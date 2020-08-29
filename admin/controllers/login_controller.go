package controllers

import "github.com/kataras/iris/context"

var loginController LoginController

type LoginController struct {
	BaseAction
}

func LoginHandler(ctx context.Context) {
	loginController.LoginAction(ctx)
}
