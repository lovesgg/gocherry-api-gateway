package middleware

import (
	"bytes"
	"fmt"
	"github.com/kataras/iris/context"
	"runtime"
)

func NewRecoverPanic() context.Handler {
	return func(ctx context.Context) {
		defer func() {
			if err := recover(); err != nil {
				if ctx.IsStopped() {
					return
				}

				var stacktrace string
				for i := 1; ; i++ {
					_, f, l, got := runtime.Caller(i)
					if !got {
						break
					}

					stacktrace += fmt.Sprintf("%s:%d\n", f, l)
				}

				var logMessage bytes.Buffer
				logMessage.WriteString(fmt.Sprintf("Recovered from a route's Handler('%s')\n", ctx.HandlerName()))
				logMessage.WriteString(fmt.Sprintf("At Request: %s\n", ctx.RemoteAddr()))
				logMessage.WriteString(fmt.Sprintf("Trace: %s\n", err))
				logMessage.WriteString(fmt.Sprintf("\n%s", stacktrace))

				ctx.StopExecution()
			}
		}()

		ctx.Next()
	}
}
