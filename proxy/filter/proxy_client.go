package filter

import (
	"context"
	"encoding/json"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"go.etcd.io/etcd/clientv3"
	"gocherry-api-gateway/components/common_enum"
	"gocherry-api-gateway/components/etcd_client"
	"gocherry-api-gateway/components/http_client"
	"gocherry-api-gateway/proxy/enum"
	"gocherry-api-gateway/proxy/model"
	"math/rand"
	"sync"
	"time"
)

//将url请求到目标ip + url 然后将结果返回
func ProxyRunToServer(proxyContext *ProxyContext, servers []string) (statusCode int, err string) {
	var responseData string
	var code int

	url := proxyContext.Url.BaseRedirectUrl //转发的地址 不是前段请求的url (有个坑 如果是get请求 前端传了参数 那么接口转发如何处理 将参数组装？)
	method := proxyContext.RequestContext.Method()
	timeRe := proxyContext.Api.TimeOut * time.Second

	if len(servers) == 0 {
		return enum.STATUS_CODE_FAILED, "网络请求节点失败"
	}

	//这是请求的完整url
	requestUrl := servers[0] + url

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

// 这里不使用服务发现 维护起来相对方便些
// 获取集群对应的全部 服务节点
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
	key := rand.Intn(len(serverList))

	return serverList[key], enum.STATUS_CODE_OK
}


/**
使用服务发现
 */
type ClientMon struct {
	client     *clientv3.Client
	serverList map[string]string
	lock       sync.Mutex
}

// 初始化server端
func NewClientMon() (*ClientMon, error) {
	return &ClientMon{
		client:     etcd_client.GetClient(),
		serverList: make(map[string]string),
	}, nil
}

func (this *ClientMon) GetService() ([]string, error) {
	key_prefix := common_enum.ETCD_KEYS_APP_SERVER_REGISTER
	resp, err := this.client.Get(context.Background(), key_prefix, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	addrs := this.extractAddrs(resp)
	go this.watcher(key_prefix)
	return addrs, nil
}

func (this *ClientMon) extractAddrs(resp *clientv3.GetResponse) []string {
	addrs := make([]string, 0)
	if resp == nil || resp.Kvs == nil {
		return addrs
	}
	for i := range resp.Kvs {
		if v := resp.Kvs[i].Value; v != nil {
			this.SetServiceList(string(resp.Kvs[i].Key), string(resp.Kvs[i].Value))
			addrs = append(addrs, string(v))
		}
	}
	return addrs
}

// watch负责将监听到的put、delete请求存放到指定list
func (this *ClientMon) watcher(prefix string) {
	rch := this.client.Watch(context.Background(), prefix, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case mvccpb.PUT:
				this.SetServiceList(string(ev.Kv.Key), string(ev.Kv.Value))
			case mvccpb.DELETE:
				this.DelServiceList(string(ev.Kv.Key))
			}
		}
	}
}

func (this *ClientMon) SetServiceList(key, val string) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.serverList[key] = string(val)
}

func (this *ClientMon) DelServiceList(key string) {
	this.lock.Lock()
	defer this.lock.Unlock()
	delete(this.serverList, key)
}
