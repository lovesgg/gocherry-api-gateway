package common_enum

const (
	REDIS_KEY_API_PROXY_CACHE = "proxy_api_" //api接口缓存前缀
	REDIS_KEY_API_PROXY_AUTH_CHECK = "proxy_auth_check_" //默认的授权插件对应的缓存token key
	REDIS_KEY_API_PROXY_LIMITING = "proxy_api_limiting_" //某个应用接口的限流

)
