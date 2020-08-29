package filter

import (
	"github.com/yangwenmai/ratelimit/simpleratelimit"
	"gocherry-api-gateway/proxy/enum"
	"time"
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
*/
func (f *LimitingFilter) Pre(proxyContext *ProxyContext) (statusCode int, err string) {
	limitNum := proxyContext.Api.LimitRequest
	if limitNum > 0 {
		rl := simpleratelimit.New(limitNum, time.Second)
		if !rl.Limit() {
			//return enum.STATUS_CODE_FAILED, "请求受限制"
		}
		//log.Printf("limit result: %v\n", rl.Limit())
		//lb := leakybucket.New()
		//_, _ = lb.Create("leaky_bucket", 10, time.Second)
	}
	return enum.STATUS_CODE_OK, "ok"
}

func (f *LimitingFilter) AfterDo(proxyContext *ProxyContext) (statusCode int, err string) {
	return enum.STATUS_CODE_OK, ""
}
