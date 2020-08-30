package controllers

import (
	"fmt"
	"github.com/kataras/iris/context"
	"gocherry-api-gateway/admin/models"
	"gocherry-api-gateway/components/utils"
)

type UserSaveReq struct {
	UserName string `json:"user_name"`
	Phone    string `json:"phone" validate:"required"`
	Pwd      string `json:"pwd" validate:"required"`
	Level    int    `json:"level" validate:"required"`
}

func (c *UserController) GetList(ctx context.Context) {
	var users []models.AdminAccount
	accounts := new(models.AdminAccount)
	list := accounts.GetUserList(1, 100, users)
	fmt.Println(list)

	c.RenderJson(ctx, list)
}

func (c *UserController) Save(ctx context.Context) {
	var req models.AdminAccount
	c.GetRequest(ctx, &req)

	var user models.AdminAccount
	user.Phone = req.Phone
	user.Pwd = utils.GetMd5(req.Pwd)
	user.Level = req.Level
	user.UserName = req.UserName
	user.State = 1

	accounts := new(models.AdminAccount)
	ret := accounts.SaveUser(user)

	c.RenderJson(ctx, ret)
}

func (c *UserController) Del(ctx context.Context) {
	var req models.AdminAccount
	c.GetRequest(ctx, &req)

	accounts := new(models.AdminAccount)
	ret := accounts.DelUser(req)

	c.RenderJson(ctx, ret)
}
