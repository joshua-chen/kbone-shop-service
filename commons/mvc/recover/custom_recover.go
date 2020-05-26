/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-25 17:50:05
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-26 18:16:01
 */

/**
* @Description: 错误信息处理
* @Author: guoyzh
* @Date: 2019/10/23
 */

package recover

import (
	"commons/mvc/models"
	"fmt"
	"runtime"
	"strconv"

	_ "github.com/kataras/golog"
	"github.com/kataras/iris/v12"
)

func CustomRecover(ctx iris.Context) {
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

			errMsg := fmt.Sprintf("错误信息: %s", err)
			// when stack finishes
			logMessage := fmt.Sprintf("从错误中恢复：('%s')\n", ctx.HandlerName())
			logMessage += errMsg + "\n"
			logMessage += fmt.Sprintf("\n%s", stacktrace)
			// 打印错误日志
			ctx.Application().Logger().Warn(logMessage)
			// 返回错误信息
			result := &models.ResponseResult{
				Msg:     errMsg,
				Success: false,
				Code:    strconv.Itoa(iris.StatusInternalServerError),
			}
			ctx.JSON(result)
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.StopExecution()
		}
	}()
	ctx.Next()
}
