package filter

import (
	"gocherry-api-gateway/components/common_enum"
	"gocherry-api-gateway/components/redis_client"
	"gocherry-api-gateway/proxy/enum"
)

type LimitingFilter struct {
	Filter
}

func (f *LimitingFilter) Init(proxyContext *ProxyContext) {
}

func (f *LimitingFilter) Name(proxyContext *ProxyContext) string {
	return Limiting
}

/**
应该是基于当前app_and_url作为唯url，基于这值做限流
目前是使用redis作为限流器
*/
func (f *LimitingFilter) Pre(proxyContext *ProxyContext) (statusCode int, err string) {
	limitNum := proxyContext.Api.LimitRequest
	if limitNum > 0 {
		key := common_enum.REDIS_KEY_API_PROXY_LIMITING + proxyContext.AppName + proxyContext.Url.Path
		redis := redis_client.GetProxyRedis()
		count := redis.Incr(key)
		if count == 1 {
			//第一次设置过期时间为1秒
			redis.Expire(key, 1)
			return enum.STATUS_CODE_OK, "ok"
		}
		if count > limitNum {
			return enum.STATUS_CODE_FAILED, "请求太激烈了"
		}
	}
	//无限制 或 没达到限制数量
	return enum.STATUS_CODE_OK, "ok"
}

func (f *LimitingFilter) AfterDo(proxyContext *ProxyContext) (statusCode int, err string) {
	return enum.STATUS_CODE_OK, ""
}
