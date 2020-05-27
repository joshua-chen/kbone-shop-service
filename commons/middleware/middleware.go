/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-27 14:28:19
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-27 18:20:05
 */
package middleware

import (
	"commons/config"
	"commons/middleware/casbins"
	"commons/middleware/jwt"
	"commons/utils"
	"strings"

	"github.com/kataras/golog"
	"github.com/kataras/iris/v12/context"

)

type Middleware struct {
}

func ServeHTTP(ctx context.Context) {
	path := ctx.Path()
	golog.Debug(path)
	// 过滤静态资源、login接口、首页等...不需要验证
	if checkURL(path) {
		ctx.Next()
		return
	}

	// jwt token拦截
	if !jwt.Serve(ctx) {
		return
	}

	// 系统菜单不进行权限拦截
	if !strings.Contains(path, "/sysMenu") {
		// casbin权限拦截
		ok := casbins.CheckPermissions(ctx)
		if !ok {
			return
		}
	}

	// Pass to real API
	ctx.Next()
}

/**
return
	true:则跳过不需验证，如登录接口等...
	false:需要进一步验证
*/
func checkURL(requestPath string) bool {
	staticPath := config.AppConfig.StaticPath
	if utils.HasPrefix(requestPath, staticPath) {
		return true
	}

	anonymousUrls := config.AppConfig.AnonymousRequset.Urls
	for _, v := range anonymousUrls {
		if requestPath == v {
			return true
		}
	}

	anonymousPrefixes := config.AppConfig.AnonymousRequset.Prefixes
	for _, v := range anonymousPrefixes {
		if utils.HasPrefix(requestPath, v) {
			return true
		}
	}

	anonymousSuffixes := config.AppConfig.AnonymousRequset.Suffixes
	for _, v := range anonymousSuffixes {
		if utils.HasSuffix(requestPath, v) {
			return true
		}
	}

	// strings.Index(requestPath,v)
	return false
}
