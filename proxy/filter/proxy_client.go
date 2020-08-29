package filter

import (
	"encoding/json"
	"gocherry-api-gateway/components/common_enum"
	"gocherry-api-gateway/components/etcd_client"
	"gocherry-api-gateway/components/http_client"
	"gocherry-api-gateway/proxy/enum"
	"gocherry-api-gateway/proxy/model"
	"time"
)

//将url请求到目标ip + url 然后将结果返回
func ProxyRunToServer(proxyContext *ProxyContext) (statusCode int, err string) {
	var responseData string
	apiData := proxyContext.Api
	url := proxyContext.Url.BaseRedirectUrl //转发的地址 不是前段请求的url (有个坑 如果是get请求 前端传了参数 那么接口转发如何处理 将参数组装？)
	method := proxyContext.RequestContext.Method()
	timeRe := proxyContext.Api.TimeOut * time.Second

	//任意取的一个服务ip 类似负载均衡 这一步需要做服务发现哦
	server, code := GetServers(proxyContext, apiData.BaseClusterName)
	if code == enum.STATUS_CODE_FAILED {
		return enum.STATUS_CODE_FAILED, "服务异常"
	}

	//这是请求的完整url
	requestUrl := server.Ip + url

	//根据不同的method发起请求 目前支持post get 其他方法暂时不支持啊 可以自行在后台定义并在这增加方法
	switch method {
	case "POST":
		params := proxyContext.RequestContext.Request().Body //获取post的数据
		decoder := json.NewDecoder(params)
		var postData map[string]string
		// 解析参数 存入map
		_ = decoder.Decode(&postData)
		responseData, code, err = http_client.Post(requestUrl, postData, timeRe)
		if code != enum.STATUS_CODE_OK {
			return enum.STATUS_CODE_FAILED, err
		}
		break
	case "GET":
		responseData, code, err = http_client.Get(requestUrl, timeRe)
		if code != enum.STATUS_CODE_OK {
			return enum.STATUS_CODE_FAILED, err
		}
		break
	default:
		//不属于这些方法禁止访问
		return enum.STATUS_CODE_FAILED, "请求方法不允许"
	}
	proxyContext.Response = responseData

	return enum.STATUS_CODE_OK, ""
}

//获取集群对应的全部 服务节点 这里使用服务发现逻辑
func GetServers(proxyContext *ProxyContext, clusterName string) (m model.Server, code int) {
	var serverList []model.Server
	var oneServer model.Server

	serverKey := common_enum.ETCD_KEYS_APP_CLUSTER_SERVER_LIST + proxyContext.AppName + "/" + clusterName
	list, _ := etcd_client.GetKvPrefix(serverKey)

	for _, value := range list.Kvs {
		_ = json.Unmarshal([]byte(value.Value), &oneServer)
		serverList = append(serverList, oneServer)
	}
	if len(serverList) == 0 {
		return oneServer, enum.STATUS_CODE_FAILED
	}
	if serverList[0].Ip == "" {
		return oneServer, enum.STATUS_CODE_FAILED
	}
	return serverList[0], enum.STATUS_CODE_OK
}
