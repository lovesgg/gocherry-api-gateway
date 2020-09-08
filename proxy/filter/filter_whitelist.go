package filter

import (
	"gocherry-api-gateway/proxy/enum"
	"strings"
)

type WhiteListFilter struct {
	Filter
}

func (f *WhiteListFilter) Init(proxyContext *ProxyContext) {
}

func (f *WhiteListFilter) Name(proxyContext *ProxyContext) string {
	return WhiteList
}

/**
只允许名单内的用户访问接口
*/
func (f *WhiteListFilter) Pre(proxyContext *ProxyContext) (statusCode int, err string) {
	whiteList := proxyContext.Api.WhiteList
	if proxyContext.Api.WhiteListCheck && whiteList != "" {
		var flag = false
		userPhone := proxyContext.UserInfo.UserPhone
		list := strings.Split(whiteList, ",")
		for _, phone := range list {
			if phone == userPhone {
				flag = true
				break
			}
		}
		//flag = false
		if flag == false {
			return enum.STATUS_CODE_FAILED, "无权限访问此接口"
		}

	}
	return enum.STATUS_CODE_OK, ""
}

func (f *WhiteListFilter) AfterDo(proxyContext *ProxyContext) (statusCode int, err string) {
	return enum.STATUS_CODE_OK, ""
}
