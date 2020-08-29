package controllers

import (
	"encoding/json"
	"github.com/kataras/iris/context"
	"gocherry-api-gateway/admin/models"
	"gocherry-api-gateway/components/common_enum"
	"gocherry-api-gateway/components/etcd_client"
	"gocherry-api-gateway/components/utils"
)

type APiSaveReq struct {
	AppName string     `json:"app_name"`
	ApiForm models.Api `json:"api_form"`
}

type ApiGetOneReq struct {
	BaseApiUrl      string `json:"base_api_url"`
	AppName         string `json:"app_name"`
	BaseClusterName string `json:"base_cluster_name"`
}

func (c *ApiController) GetList(ctx context.Context) {
	var req APiSaveReq
	var apiList []models.Api
	var oneApi models.Api

	c.GetRequest(ctx, &req)

	apiKey := common_enum.ETCD_KEYS_APP_API_LIST + req.AppName
	//etcd_client.DelKvPrefix(apiKey) //谨慎使用
	list, _ := etcd_client.GetKvPrefix(apiKey)

	for _, value := range list.Kvs {
		_ = json.Unmarshal([]byte(value.Value), &oneApi)
		apiList = append(apiList, oneApi)
	}
	if len(apiList) > 0 {
		c.RenderJson(ctx, apiList)
	} else {
		c.RenderError(ctx, common_enum.ComError{Msg: "无数据"})
	}
}

func (c *ApiController) Save(ctx context.Context) {
	var req APiSaveReq
	c.GetRequest(ctx, &req)

	apiKey := common_enum.ETCD_KEYS_APP_API_LIST + req.AppName + req.ApiForm.BaseApiUrl
	appExists, _ := etcd_client.GetKv(apiKey)
	if appExists != nil {

	}
	req.ApiForm.UpdateTime = utils.GetNowTimeFormat()

	newApi, _ := json.Marshal(req.ApiForm)
	_, _ = etcd_client.PutKv(apiKey, string(newApi))

	c.RenderJson(ctx, "api创建成功")
}

func (c *ApiController) GetOne(ctx context.Context) {
	var req ApiGetOneReq
	var oneApi models.Api
	c.GetRequest(ctx, &req)

	apiKey := common_enum.ETCD_KEYS_APP_API_LIST + req.AppName + req.BaseApiUrl
	data, _ := etcd_client.GetKvPrefix(apiKey)
	_ = json.Unmarshal([]byte(data.Kvs[0].Value), &oneApi)

	c.RenderJson(ctx, oneApi)

}

func (c *ApiController) Del(ctx context.Context) {
	var req ApiGetOneReq
	c.GetRequest(ctx, &req)
	apiKey := common_enum.ETCD_KEYS_APP_API_LIST + req.AppName + req.BaseApiUrl
	ret, _ := etcd_client.DelKv(apiKey)
	if ret != nil {

	}

	c.RenderJson(ctx, "删除成功")
}