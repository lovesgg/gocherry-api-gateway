package filter

import (
	"github.com/kataras/iris/context"
	"gocherry-api-gateway/components/log_client"
	"gocherry-api-gateway/proxy/enum"
	"time"
)

/**
对于接口的前缀校验 比如状态 请求方法等 只有走通了才会走后续的流程
特别需要注意 建议先走这插件再走其他流程
*/

type PreFilter struct {
	Filter
}

func (f *PreFilter) Init(proxyContext *ProxyContext) {
}

func (f *PreFilter) Name(proxyContext *ProxyContext) string {
	return Pre
}

func (f *PreFilter) Pre(proxyContext *ProxyContext) (statusCode int, err string) {
	nowTime := time.Now().Unix()
	proxyContext.StartTime = nowTime
	//记录接口开始日志
	log_client.LogInfo(proxyContext.RequestContext, map[string]interface{}{
		"time": nowTime,
		"app_name": proxyContext.AppName,
		"url": proxyContext.Url.Path,
		"redirect_url": proxyContext.Url.BaseRedirectUrl,
		"params": "",//参数
	})

	//判断接口是否被禁用
	apiStatus := proxyContext.Api.BaseApiStatus
	if apiStatus == false {
		return enum.STATUS_CODE_FAILED, "接口无法访问"
	}

	requestMethods := proxyContext.Api.BaseApiRequestType //允许请求的方法
	method := proxyContext.RequestContext.Method()

	//判断请求方法是否允许
	methodFlag := false
	for _, value := range requestMethods {
		if value == method {
			methodFlag = true //只要有一个方法允许即可
			break
		}
	}
	if methodFlag == false {
		return enum.STATUS_CODE_FAILED, "请求方法不允许"
	}

	//判断接口是否降级
	reduceLevel := proxyContext.Api.ReduceLevel
	if reduceLevel {
		return enum.STATUS_CODE_FAILED, "接口暂时无法访问"
	}

	return enum.STATUS_CODE_OK, ""
}

func (f *PreFilter) AfterDo(proxyContext *ProxyContext) (statusCode int, err string) {
	return enum.STATUS_CODE_OK, ""
}

/**
从头部获取app_name
*/
func GetAppName(ctx context.Context) string {
	appName := ctx.GetHeader("app_name")
	return appName
}
