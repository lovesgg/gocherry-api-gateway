package main

import (
	"encoding/json"
	"fmt"
	"gocherry-api-gateway/admin/controllers"
	"gocherry-api-gateway/components/common_enum"
	"gocherry-api-gateway/components/etcd_client"
	"gocherry-api-gateway/components/http_client"
	"gocherry-api-gateway/proxy/enum"
	"sync"
	"time"
)

/**
每5分钟校验全部ip 如果访问不到就删除此server节点
*/

func main() {
	var oneServer controllers.ServerSaveReq
	var serverList []controllers.ServerSaveReq
	for {
		serverKey := common_enum.ETCD_KEYS_APP_CLUSTER_SERVER_LIST
		list, _ := etcd_client.GetKvPrefix(serverKey)

		for _, value := range list.Kvs {
			_ = json.Unmarshal([]byte(value.Value), &oneServer)
			serverList = append(serverList, oneServer)
		}

		length := len(serverList)
		if length > 0 {
			wg := sync.WaitGroup{}
			wg.Add(length)
			for _, ip := range serverList {
				go func(ip controllers.ServerSaveReq) {
					res, code, msg := http_client.Get(ip.Ip+"/health/check", 1 * time.Second)
					if code != enum.STATUS_CODE_OK {
						//在这禁用 禁用前根据实际业务需要比如请求失败5次就删除节点
						fmt.Println(ip.Ip, res, code, msg, "禁用server")
					}
					wg.Done()
				}(ip)
			}
			wg.Wait()
		}

		time.Sleep(300 * time.Second)
	}
}
