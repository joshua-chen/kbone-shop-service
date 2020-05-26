/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-22 15:53:19
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-26 18:37:27
 */
package mvc

import (
	_ "commons/config"
	_ "commons/middleware/cors"
	_ "commons/middleware/jwt"

	_ "github.com/dgrijalva/jwt-go/request"
	_ "github.com/iris-contrib/middleware/cors"
	_ "github.com/iris-contrib/swagger/v12"
	_ "github.com/iris-contrib/swagger/v12/swaggerFiles"
	_ "github.com/kataras/golog"
	_ "github.com/kataras/iris/v12"
	_ "github.com/kataras/iris/v12/context"
	_ "github.com/kataras/iris/v12/core/router"
	_ "github.com/kataras/iris/v12/hero"
	_ "github.com/kataras/iris/v12/i18n"
	_ "github.com/kataras/iris/v12/middleware/logger"
	_ "github.com/kataras/iris/v12/middleware/recover"
)

 