/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-27 14:28:19
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-28 15:33:40
 */
package auth

import (
	"commons/config"
	"commons/middleware/casbin"
	"commons/middleware/jwt"
	"commons/utils"
	"strings"
	"sync"

	"github.com/kataras/golog"
	"github.com/kataras/iris/v12/context"

)

type Auth struct {
}

var (
	instance *Auth
	lock     *sync.Mutex = &sync.Mutex{}
)

func Instance() *Auth {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &Auth{}
		}
	}
	return instance
}
func (a *Auth) New() context.Handler {

	return New()
}

func New() context.Handler {
	return func(ctx context.Context) {
		path := ctx.Path()
		golog.Debug(path)
		// 过滤静态资源、login接口、首页等...不需要验证
		if checkURL(path) {
			ctx.Next()
			return
		}

		// jwt token拦截
		if !jwt.Filter(ctx) {
			return
		}

		// 系统菜单不进行权限拦截
		if !strings.Contains(path, "/sysMenu") {
			// casbin权限拦截
			ok := casbin.Filter(ctx)
			if !ok {
				return
			}
		}

		// Pass to real API
		ctx.Next()
	}
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
