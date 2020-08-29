package controllers

import (
	"encoding/json"
	"github.com/kataras/iris/context"
	"gocherry-api-gateway/components/common_enum"
	"gocherry-api-gateway/components/etcd_client"
	"gocherry-api-gateway/components/utils"
)

type ServerSaveReq struct {
	ClusterName string `json:"cluster_name" validate:"required"`
	AppName     string `json:"app_name" validate:"required"`
	ServerName  string `json:"server_name"` //随机取一个不重复的id值 时间戳也可以 不让前端传过来了
	Ip          string `json:"ip"`          //可以是ip或者是http开头的 不做严格限制
	UpdateTime  string `json:"update_time"`
}

func (c *ServerController) GetList(ctx context.Context) {
	var req ServerSaveReq
	c.GetRequest(ctx, &req)

	var serverList []ServerSaveReq
	var oneServer ServerSaveReq

	serverKey := common_enum.ETCD_KEYS_APP_CLUSTER_SERVER_LIST + req.AppName + "/" + req.ClusterName
	list, _ := etcd_client.GetKvPrefix(serverKey)

	for _, value := range list.Kvs {
		_ = json.Unmarshal([]byte(value.Value), &oneServer)
		serverList = append(serverList, oneServer)
	}
	if len(serverList) > 0 {
		c.RenderJson(ctx, serverList)
	} else {
		c.RenderError(ctx, common_enum.ComError{Msg: "无数据"})
	}
}

func (c *ServerController) Save(ctx context.Context) {
	var req ServerSaveReq
	c.GetRequest(ctx, &req)

	serverKey := common_enum.ETCD_KEYS_APP_CLUSTER_SERVER_LIST + req.AppName + "/" + req.ClusterName + "/" + req.Ip
	appExists, _ := etcd_client.GetKv(serverKey)
	if appExists != nil {

	}
	req.UpdateTime = utils.GetNowTimeFormat()

	newServer, _ := json.Marshal(req)
	_, _ = etcd_client.PutKv(serverKey, string(newServer))

	c.RenderJson(ctx, "节点创建成功")
}
