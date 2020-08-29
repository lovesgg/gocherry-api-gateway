package filter

import (
	"gocherry-api-gateway/proxy/enum"
	"strings"
)

type BlackListFilter struct {
	Filter
}

func (f *BlackListFilter) Init(proxyContext *ProxyContext) {

}

func (f *BlackListFilter) Name(proxyContext *ProxyContext) string {
	return BlackList
}

//校验ip是否在黑名单列表
func (f *BlackListFilter) Pre(proxyContext *ProxyContext) (statusCode int, err string) {
	blackList := proxyContext.Api.IpBlack
	userIp := proxyContext.RequestContext.Request().RemoteAddr
	userIps := strings.Split(userIp, ":")
	if blackList != "" {
		lists := strings.Split(blackList, ",")
		for _, ip := range lists {
			if ip == userIps[0] {
				return enum.STATUS_CODE_FAILED, "您已被禁用"
			}
		}
	}
	return enum.STATUS_CODE_OK, ""
}

func (f *BlackListFilter) AfterDo(proxyContext *ProxyContext) (statusCode int, err string) {
	return enum.STATUS_CODE_OK, ""
}
