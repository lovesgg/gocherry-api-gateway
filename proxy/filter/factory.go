package filter

///**
//插件公共接口 其他插件都继承于这插件去开发
//*/
type Filter interface {
	Init(ctx *ProxyContext) //初始方法
	Name(ctx *ProxyContext) string //返回插件名
	Pre(ctx *ProxyContext) (statusCode int, err string) //请求下游接口之前执行的方法
	AfterDo(ctx *ProxyContext) (statusCode int, er string) //请求完成下游接口后的方法
}

type BaseFilter struct {
	fList map[string]Filter //所有的插件都放在这里 遍历执行即可
}

func (f *BaseFilter) init() {
	f.fList = make(map[string]Filter)
}

func (f *BaseFilter) register(name string, filter Filter) {
	f.fList[name] = filter
}
