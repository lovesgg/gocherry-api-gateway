package filter

import "gocherry-api-gateway/proxy/enum"

type JwtFilter struct {
	Filter
}

func (f *JwtFilter) Init(proxyContext *ProxyContext) {
}

func (f *JwtFilter) Name(proxyContext *ProxyContext) string {
	return Jwt
}

func (f *JwtFilter) Pre(proxyContext *ProxyContext) (statusCode int, err string) {
	return enum.STATUS_CODE_OK, ""
}

func (f *JwtFilter) AfterDo(proxyContext *ProxyContext) (statusCode int, err string) {
	return enum.STATUS_CODE_OK, ""
}