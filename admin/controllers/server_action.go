package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	irisContext "github.com/kataras/iris/context"
	"go.etcd.io/etcd/clientv3"
	"gocherry-api-gateway/components/common_enum"
	"gocherry-api-gateway/components/etcd_client"
	"gocherry-api-gateway/components/utils"
	"time"
)

type ServerSaveReq struct {
	ClusterName string `json:"cluster_name" validate:"required"`
	AppName     string `json:"app_name" validate:"required"`
	ServerName  string `json:"server_name"` //随机取一个不重复的id值 时间戳也可以 不让前端传过来了
	Ip          string `json:"ip"`          //可以是ip或者是http开头的 不做严格限制
	UpdateTime  string `json:"update_time"`
}

func (c *ServerController) GetList(ctx irisContext.Context) {
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

func (c *ServerController) Save(ctx irisContext.Context) {
	var req ServerSaveReq
	c.GetRequest(ctx, &req)

	serverKey := common_enum.ETCD_KEYS_APP_CLUSTER_SERVER_LIST + req.AppName + "/" + req.ClusterName + "/" + req.Ip
	appExists, _ := etcd_client.GetKv(serverKey)
	if appExists != nil {

	}
	req.UpdateTime = utils.GetNowTimeFormat()

	_ = RegisterServer(req)

	//存储详情 用于列表展示 不作为服务发现
	newServer, _ := json.Marshal(req)
	_, _ = etcd_client.PutKv(serverKey, string(newServer))

	c.RenderJson(ctx, "节点创建成功")
}

func (c *ServerController) Del(ctx irisContext.Context) {
	var req ServerSaveReq
	c.GetRequest(ctx, &req)
	serverKey := common_enum.ETCD_KEYS_APP_CLUSTER_SERVER_LIST + req.AppName + "/" + req.ClusterName + "/" + req.Ip
	ret, _ := etcd_client.DelKv(serverKey)
	if ret != nil {

	}

	c.RenderJson(ctx, "节点删除成功")
}

/**
服务注册
*/

//注册服务
func RegisterServer(requestData ServerSaveReq) error {
	client := etcd_client.GetClient()
	kv := clientv3.NewKV(client)
	key_prefix := common_enum.ETCD_KEYS_APP_SERVER_REGISTER + requestData.Ip
	ctx := context.Background()
	lease := clientv3.NewLease(client)

	leaseRes, err := clientv3.NewLease(client).Grant(ctx, 300)
	if err != nil {
		return err
	}
	_, err = kv.Put(context.Background(), key_prefix, requestData.Ip, clientv3.WithLease(leaseRes.ID)) //把服务的key绑定到租约下面
	if err != nil {
		return err
	}
	//续租时间大概自动为租约的三分之一时间，context.TODO官方定义为是你不知道要传什么
	keepaliveRes, err := lease.KeepAlive(context.TODO(), leaseRes.ID)
	if err != nil {
		return err
	}
	go lisKeepAlive(keepaliveRes)
	return err
}

/**
续约
*/
func lisKeepAlive(keepaliveRes <-chan *clientv3.LeaseKeepAliveResponse) {
	for {
		select {
		case ret := <-keepaliveRes:
			if ret != nil {
				fmt.Println("续租成功", time.Now())
			}
		}
	}
}