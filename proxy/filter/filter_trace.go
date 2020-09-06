package filter

import (
	"gocherry-api-gateway/components/log_client"
	"gocherry-api-gateway/proxy/enum"
	"math/rand"
	"time"
)

/**
记录日志id和记录时间
*/

type TraceFilter struct {
	Filter
}

func (f *TraceFilter) Init(proxyContext *ProxyContext) {
}

func (f *TraceFilter) Name(proxyContext *ProxyContext) string {
	return TraceF
}

//设置trace_id相关 放到header传到下游
func (f *TraceFilter) Pre(proxyContext *ProxyContext) (statusCode int, err string) {
	nowTime := time.Now().Unix()
	//trace_id 应该需要记录手机号或者用户id 好排查问题
	proxyContext.TraceId = int64(rand.Intn(500)) + nowTime + 2 + int64(rand.Intn(1000))

	return enum.STATUS_CODE_OK, ""
}

func (f *TraceFilter) AfterDo(proxyContext *ProxyContext) (statusCode int, err string) {
	proxyContext.EndTime = time.Now().Unix()

	log_client.LogInfo(proxyContext.RequestContext, map[string]interface{}{
		"start_time": proxyContext.StartTime,
		"end_time":   proxyContext.EndTime,
		"log_time":   time.Now().Unix() - proxyContext.StartTime,
		"trace_id":   proxyContext.TraceId,
		"url":        proxyContext.Url.Path,
		"app_name":   proxyContext.AppName,
	})

	return enum.STATUS_CODE_OK, ""
}
