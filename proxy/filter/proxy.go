package filter

import (
	"encoding/json"
	"github.com/kataras/iris/context"
	simpleRate "github.com/yangwenmai/ratelimit/simpleratelimit"
	"gocherry-api-gateway/admin/controllers"
	"gocherry-api-gateway/admin/models"
	"gocherry-api-gateway/components/common_enum"
	"gocherry-api-gateway/components/etcd_client"
	"gocherry-api-gateway/proxy/enum"
	"sync"
)

type ProxyController struct {
	controllers.BaseAction
}

/**
完整生命周期都需要携带的上下文
*/
type ProxyContext struct {
	AppName        string
	Url            Url
	Api            models.Api
	Mutex          sync.RWMutex
	AppConfig      *common_enum.Config
	Filters        BaseFilter
	RequestContext context.Context
	baseController *ProxyController
	Response       string
	StartTime      int64
	EndTime        int64
	TraceId        int64
	RateLimiter    *simpleRate.RateLimiter
	UserInfo       UserInfo
}

type Url struct {
	Path            string
	FullPath        string
	BaseRedirectUrl string
	AppAndUrl       string //多应用全局唯一url 缓存key 等都需要使用这key
}

type UserInfo struct {
	UserId    string
	UserPhone string
	UserToken string
	UserName  string
}

/**
接口唯一入口url处理 核心处理入口
*/
func (c *ProxyController) DoProxyHandler(ctx context.Context, appConfig *common_enum.Config, servers *ClientMon) {
	var proxyContext *ProxyContext
	//首先验证app_name
	appName := GetAppName(ctx)
	if appName == "" {
		c.RenderError(ctx, common_enum.ComError{Msg: "app_name 必填"})
		return
	}

	proxyContext = GetProxyContext(ctx, c, appConfig, appName)

	//初始化插件
	filters := InitFilters()
	//遍历执行插件->请求前的插件 只执行pre方法
	statusCode, err := RunFilters(filters, proxyContext)
	if statusCode != enum.STATUS_CODE_OK {
		if statusCode == enum.STATUS_CODE_FAILED {
			//出问题统一使用这方法输出
			c.RenderError(ctx, common_enum.ComError{Msg: err})
			return
		}
		if statusCode == enum.STATUS_CODE_HAS_CACHE {
			//直接输出缓存 结束（不再走后续的接口后缀插件啦）
			_, _ = proxyContext.RequestContext.WriteString(proxyContext.Response)
			return
		}
		if statusCode == enum.STATUS_CODE_AUTH_NOT_LOGIN {
			//直接输出缓存 结束（不再走后续的接口后缀插件啦）
			c.RenderError(ctx, common_enum.ComError{Msg: "未登录", Code: 1024})
			return
		}
	}

	//初始化插件没问题后组装请求到目标服务对应的地址，将结果返回
	res, msg := ProxyRunToServer(proxyContext, servers)
	if res != enum.STATUS_CODE_OK {
		c.RenderError(ctx, common_enum.ComError{Msg: msg})
		return
	}

	//遍历执行插件->请求后的插件 只执行after方法
	//这里只作为请求后的辅助插件 比如放缓存 记日志啥的 不建议放核心插件
	afterCode, err := RunFiltersAfter(filters, proxyContext)
	if afterCode != enum.STATUS_CODE_OK {
		if afterCode == enum.STATUS_CODE_FAILED {
			c.RenderError(ctx, common_enum.ComError{Msg: err})
			return
		}
	}

	//输出下游接口的请求结果
	_, _ = proxyContext.RequestContext.WriteString(proxyContext.Response)
	return
}

func GetUrl(ctx context.Context) string {
	return ctx.Path()
}

func GetFullUrl(ctx context.Context) string {
	return ctx.Request().RequestURI
}

//组装上下文数据 只有返回此数据后才可使用ProxyContext 不要在这方法内部调用其他方法时乱使用ProxyContext的数据 因为有的数据还没组装好
func GetProxyContext(ctx context.Context, controller *ProxyController, appConfig *common_enum.Config, appName string) *ProxyContext {
	var proxyContext = new(ProxyContext)

	url := GetUrl(ctx)                  //不带参数的url 用这个去获取etcd的api数据
	fullUrl := GetFullUrl(ctx)          //完整url 用这个url去请求下游
	apiData := GetApiData(appName, url) //接口配置信息

	proxyContext.Url.Path = url
	proxyContext.Url.FullPath = fullUrl
	proxyContext.Url.BaseRedirectUrl = apiData.BaseRedirectUrl
	proxyContext.Url.AppAndUrl = appName + url
	proxyContext.RequestContext = ctx
	proxyContext.baseController = controller
	proxyContext.AppConfig = appConfig
	proxyContext.AppName = appName
	proxyContext.Api = apiData

	return proxyContext
}

//获取api接口配置
func GetApiData(appName string, url string) models.Api {
	var apiModel models.Api

	apiKey := common_enum.ETCD_KEYS_APP_API_LIST + appName + url
	data, _ := etcd_client.GetKvPrefix(apiKey)
	_ = json.Unmarshal([]byte(data.Kvs[0].Value), &apiModel)
	return apiModel
}
