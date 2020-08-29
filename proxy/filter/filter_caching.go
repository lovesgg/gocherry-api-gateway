package filter

import (
	"gocherry-api-gateway/components/common_enum"
	"gocherry-api-gateway/components/redis_client"
	"gocherry-api-gateway/proxy/enum"
	"time"
)

type CachingFilter struct {
	Filter
}

func (f *CachingFilter) Init(proxyContext *ProxyContext) {
}

func (f *CachingFilter) Name(proxyContext *ProxyContext) string {
	return Caching
}

//读取redis缓存
func (f *CachingFilter) Pre(proxyContext *ProxyContext) (statusCode int, err string) {
	expireTime := proxyContext.Api.CacheSave
	if expireTime <= 0 {
		//缓存不会有数据
		return enum.STATUS_CODE_OK, ""
	}
	redis := redis_client.GetProxyRedis()
	data := redis.Get(getKey(proxyContext))
	proxyContext.Response = data
	if data != "" {
		return enum.STATUS_CODE_HAS_CACHE, ""
	} else {
		return enum.STATUS_CODE_OK, ""
	}
}

//set 缓存
func (f *CachingFilter) AfterDo(proxyContext *ProxyContext) (statusCode int, err string) {
	//校验缓存时间配置
	expireTime := proxyContext.Api.CacheSave
	if expireTime <= 0 {
		//不放缓存
		return enum.STATUS_CODE_OK, ""
	}

	redis := redis_client.GetProxyRedis()
	redis.Set(getKey(proxyContext), proxyContext.Response, expireTime*time.Second)

	return enum.STATUS_CODE_OK, ""
}

func getKey(proxyContext *ProxyContext) string {
	return common_enum.REDIS_KEY_API_PROXY_CACHE + proxyContext.Url.AppAndUrl
}
