package controllers

import (
	"encoding/json"
	"github.com/kataras/iris/context"
	"gocherry-api-gateway/components/common_enum"
	"gocherry-api-gateway/components/etcd_client"
	"gocherry-api-gateway/components/utils"
)

type ClusterSaveReq struct {
	ClusterName string `json:"cluster_name"`
	Title       string `json:"title"`
	Detail      string `json:"detail"`
	AppName     string `json:"app_name" validate:"required"`
	UpdateTime  string `json:"update_time"`
}

func (c *ClusterController) GetList(ctx context.Context) {
	var req ClusterSaveReq
	c.GetRequest(ctx, &req)

	var clusterList []ClusterSaveReq
	var oneCluster ClusterSaveReq

	clusterKey := common_enum.ETCD_KEYS_APP_CLUSTER_LIST + req.AppName
	list, _ := etcd_client.GetKvPrefix(clusterKey)

	for _, value := range list.Kvs {
		_ = json.Unmarshal([]byte(value.Value), &oneCluster)
		clusterList = append(clusterList, oneCluster)
	}
	if len(clusterList) > 0 {
		c.RenderJson(ctx, clusterList)
	} else {
		c.RenderError(ctx, common_enum.ComError{Msg: "无数据"})
	}
}

func (c *ClusterController) Save(ctx context.Context) {
	var req ClusterSaveReq
	c.GetRequest(ctx, &req)

	appKey := common_enum.ETCD_KEYS_APP_CLUSTER_LIST + req.AppName + "/" + req.ClusterName
	appExists, _ := etcd_client.GetKv(appKey)
	if appExists != nil {

	}
	req.UpdateTime = utils.GetNowTimeFormat()

	newCluster, _ := json.Marshal(req)
	_, _ = etcd_client.PutKv(appKey, string(newCluster))

	c.RenderJson(ctx, "集群创建成功")
}
