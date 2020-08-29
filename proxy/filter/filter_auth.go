package filter

import (
	"fmt"
	"gocherry-api-gateway/components/common_enum"
	"gocherry-api-gateway/components/redis_client"
	"gocherry-api-gateway/proxy/enum"
)

/**
默认的授权插件 这里根据自己业务实际需要做开发:如先从缓存里面读取用户登录信息 取不到返回前端重新发起登录(然后再放入缓存)
*/
type AuthFilter struct {
	Filter
}

func (f *AuthFilter) Init(proxyContext *ProxyContext) {

}

func (f *AuthFilter) Name(proxyContext *ProxyContext) string {
	return AuthF
}

func (f *AuthFilter) Pre(proxyContext *ProxyContext) (statusCode int, err string) {
	if proxyContext.Api.UserAuth {
		redis := redis_client.GetProxyRedis()
		token := redis.Get(GetAuthCacheKey(proxyContext))
		if token == "" {
			//需要由前端重新发起登录
			return enum.STATUS_CODE_AUTH_NOT_LOGIN, "未登录"
		}
		//如果已经登录 需要把用户信息状态放在ctx

	}

	return enum.STATUS_CODE_OK, ""
}

func (f *AuthFilter) AfterDo(proxyContext *ProxyContext) (statusCode int, err string) {
	return enum.STATUS_CODE_OK, ""
}

func GetAuthCacheKey(proxyContext *ProxyContext) string {
	header := proxyContext.RequestContext.Request().Header["auth_token"]
	fmt.Println("hhh", header)
	key := common_enum.REDIS_KEY_API_PROXY_AUTH_CHECK + proxyContext.AppName + "_" + string("l")
	return key
}
