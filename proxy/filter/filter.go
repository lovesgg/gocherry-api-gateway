package filter

import (
	"fmt"
	"gocherry-api-gateway/proxy/enum"
)

const (
	AuthF     = "AuthFilter"
	BlackList = "BlackListFilter"
	Caching   = "CachingFilter"
	Jwt       = "JwtFilter"
	Limiting  = "LimitingFilter"
	TraceF    = "TraceFilter"
	Pre       = "PreFilter"
	WhiteList = "WhiteList"
)

/**
初始化插件
*/
func InitFilters() BaseFilter {
	var baseFilter BaseFilter

	//初始化插件 数组
	baseFilter.init()

	//配置数据 使得可以遍历 -> 按优先级遍历
	//最好是把关键的必要插件放在最前面 免得后续的插件走不通
	allFilters := map[string]Filter{
		Pre:       new(PreFilter),
		BlackList: new(BlackListFilter),
		WhiteList: new(WhiteListFilter),
		Limiting:  new(LimitingFilter),
		AuthF:     new(AuthFilter),
		Caching:   new(CachingFilter),
		Jwt:       new(JwtFilter),
		TraceF:    new(TraceFilter),
	}

	for filterName, filter := range allFilters {
		baseFilter.register(filterName, filter)
	}

	return baseFilter
}

//遍历执行插件->请求前的
func RunFilters(filters BaseFilter, ctx *ProxyContext) (statusCode int, err string) {
	for _, filter := range filters.fList {
		filter.Init(ctx)
		filterName := filter.Name(ctx)
		statusCode, errMsg := filter.Pre(ctx)
		fmt.Println("run filters before: ", filterName, statusCode)

		if statusCode != enum.STATUS_CODE_OK {
			if statusCode == enum.STATUS_CODE_FAILED {
				fmt.Println(filterName+" before run error exit:", errMsg)
				return statusCode, errMsg
			} else if statusCode == enum.STATUS_CODE_HAS_CACHE {
				return statusCode, errMsg
			}
			return statusCode, errMsg
		}
	}
	return enum.STATUS_CODE_OK, ""
}

//遍历执行插件-》请求后的
//如果取的是缓存数据 就不走这里啦 所以后缀插件不要放核心逻辑 避免执行不到哦 后缀插件放点小功能即可呢
func RunFiltersAfter(filters BaseFilter, ctx *ProxyContext) (statusCode int, err string) {
	for _, filter := range filters.fList {
		filter.Init(ctx)
		filterName := filter.Name(ctx)
		statusCode, errMsg := filter.AfterDo(ctx)
		fmt.Println("run filters after: ", filterName, statusCode)

		if statusCode != enum.STATUS_CODE_OK {
			fmt.Println(filterName+" after run error exit:", errMsg)
			return statusCode, errMsg
		}
	}
	return enum.STATUS_CODE_OK, ""
}
