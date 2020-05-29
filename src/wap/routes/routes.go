/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-16 23:39:39
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-28 21:16:22
 */

package routes

import (
	_"commons/middleware"
	_"commons/middleware/jwt"
	_ "commons/mvc/route"
	_ "commons/mvc/context/response"
	_ "fmt"
	_ "net/http"
	_"shop/repositories"
	_"shop/services"
	"shop/wap/routes/products"
	_ "strings"

	_ "github.com/jmespath/go-jmespath"
	"github.com/kataras/iris/v12"
	_ "github.com/kataras/iris/v12/context"
	_"github.com/kataras/iris/v12/hero"
	_ "github.com/kataras/iris/v12/mvc"

)


 
func Register(app *iris.Application) {
	products.Register(app)
}


