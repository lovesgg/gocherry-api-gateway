package log_client

import (
	"github.com/kataras/iris/context"
	"gocherry-api-gateway/admin/services"
	"time"
)

var logFileName string

func init() {
	x := time.Date(2017, 02, 27, 17, 30, 20, 20, time.Local)
	logDir := services.GetAppConfig().Common.LogDir
	logFileName = logDir + "proxy_log_" + x.Format("2006-01-02") + ".log"
}

func LogInfo(ctx context.Context, log interface{}) {
	ctx.Application().Logger().Info(log)
}

func LogErr(ctx context.Context, log interface{}) {
	ctx.Application().Logger().Error(log)
}

func logWarn(ctx context.Context, log interface{}) {
	ctx.Application().Logger().Warn(log)
}
