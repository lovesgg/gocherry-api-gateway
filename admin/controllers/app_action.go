package controllers

import (
	"encoding/json"
	"github.com/kataras/iris/context"
	"gocherry-api-gateway/components/common_enum"
	"gocherry-api-gateway/components/etcd_client"
	"gocherry-api-gateway/components/utils"
)

type AppSaveReq struct {
	AppName    string `json:"app_name"`
	Title      string `json:"title"`
	Detail     string `json:"detail"`
	UpdateTime string `json:"update_time"`
}

func (c *AppController) GetList(ctx context.Context) {
	var appList []AppSaveReq
	var oneApp AppSaveReq

	appKey := common_enum.ETCD_KEYS_APP_LIST
	list, _ := etcd_client.GetKvPrefix(appKey)

	for _, value := range list.Kvs {
		_ = json.Unmarshal([]byte(value.Value), &oneApp)
		appList = append(appList, oneApp)
	}
	if len(appList) > 0 {
		c.RenderJson(ctx, appList)
	} else {
		c.RenderError(ctx, common_enum.ComError{Msg: "无数据"})
	}
}

func (c *AppController) Save(ctx context.Context) {
	var req AppSaveReq
	c.GetRequest(ctx, &req)

	appKey := common_enum.ETCD_KEYS_APP_LIST + req.AppName
	appExists, _ := etcd_client.GetKv(appKey)
	if appExists != nil {

	}
	req.UpdateTime = utils.GetNowTimeFormat()

	newApp, _ := json.Marshal(req)
	_, _ = etcd_client.PutKv(appKey, string(newApp))

	c.RenderJson(ctx, "应用创建成功")
}
