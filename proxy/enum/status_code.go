package enum

//建议统一使用 STATUS_CODE_OK 作为唯一合法状态 其余的状态都需要做下 if 校验
const (
	STATUS_CODE_OK     = 1 //成功标识
	STATUS_CODE_FAILED = 0 //失败标识
	STATUS_CODE_HAS_CACHE = 2 //走缓存并且取到缓存标识
	STATUS_CODE_AUTH_NOT_LOGIN = 3 //未登录
)
