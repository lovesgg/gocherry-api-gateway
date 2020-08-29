package controllers

import (
	"github.com/kataras/iris/context"
	"gocherry-api-gateway/components/common_enum"
	"gopkg.in/go-playground/validator.v9"
)

type BaseAction struct {
}

//返回错误，并且打warn日志
func (BaseAction) RenderError(ctx context.Context, error common_enum.ComError) {
	_, _ = ctx.JSON(map[string]interface{}{
		"ret": 0,
		"error": map[string]interface{}{
			"msg":  error.Msg,
			"code": error.Code,
		},
	})
}

func (BaseAction) RenderJson(ctx context.Context, data interface{}) {
	_, _ = ctx.JSON(map[string]interface{}{
		"ret":  1,
		"data": data,
	})
}

func (BaseAction) RenderJsonData(ctx context.Context, data interface{}, error common_enum.ComError) {
	_, _ = ctx.JSON(map[string]interface{}{
		"ret":  1,
		"data": data,
		"error_msg": map[string]interface{}{
			"msg":  error.Msg,
			"code": error.Code,
		},
	})
}

// 获取请求数据对象
func (BaseAction) GetRequest(ctx context.Context, req interface{}) {
	if err := ctx.ReadJSON(req); err != nil {

	}

	v := validator.New()
	errV := v.Struct(req)
	if errV != nil {

	}
}
