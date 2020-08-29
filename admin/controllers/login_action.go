package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/context"
	"gocherry-api-gateway/admin/models"
	"gocherry-api-gateway/components/common_enum"
	"gocherry-api-gateway/components/etcd_client"
	"gocherry-api-gateway/components/utils"
)

type loginReq struct {
	Password string `json:"password" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
}

func (c *LoginController) LoginAction(ctx context.Context) {
	var req loginReq
	var user models.AdminAccount
	var appList []AppSaveReq
	var oneApp AppSaveReq
	c.GetRequest(ctx, &req)

	//登录验证
	admin := new(models.AdminAccount)
	user = admin.GetUserBydPwd(req.Phone, utils.GetMd5(req.Password), user)

	appKey := common_enum.ETCD_KEYS_APP_LIST
	list, _ := etcd_client.GetKvPrefix(appKey)

	for _, value := range list.Kvs {
		_ = json.Unmarshal([]byte(value.Value), &oneApp)
		appList = append(appList, oneApp)
	}
	fmt.Println(user.Id, user.ID)
	if user.Phone != "" {
		c.RenderJson(ctx, map[string]interface{}{
			"login_token": req.Phone,
			"app_list":    appList,
		})
	} else {
		err := common_enum.ComError{
			Msg: "账号或密码错误",
		}
		c.RenderError(ctx, err)
	}
}
