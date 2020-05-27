/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-16 23:39:39
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-27 16:58:50
 */

package routes

import (
	"commons/middleware"
	_"commons/middleware/jwt"
	_ "commons/middleware/jwt/route"
	_ "commons/mvc/context/response"
	_ "fmt"
	_ "net/http"
	"shop/repositories"
	"shop/services"
	_ "strings"

	_ "github.com/jmespath/go-jmespath"
	"github.com/kataras/iris/v12"
	_ "github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/hero"
	_ "github.com/kataras/iris/v12/mvc"

)

var (
	BaseUri = "/api/v1"
)

/*
$ go get github.com/jmespath/go-jmespath
*/
func NewApp() *iris.Application {
	app := iris.New()
	// PartyFunc 等同于 usersRouter := app.Party("/users")
	//但它为我们提供了一种简单的方法来调用路由组的注册路由方法，
	//即来自另一个可以处理这组路由的包的函数。
	return app
}

func Register(app *iris.Application) {
	//jwtServe := jwt.GetJWT().Serve;	
	//app.Use(jwt.GetJWT().Serve) // jwt
	api := app.Party(BaseUri)	
	api.Use(middleware.ServeHTTP)
	//api.PartyFunc("/{name:string}", registerProductRoutes)
	registerProducts(api)
}

/*
开始使用路由
*/
func registerProducts(router iris.Party) {
	deps := hero.New()
	repo := repositories.NewProductRepository()
	productService := services.NewProductService(repo)
	deps.Register(productService)
	
	// GET: /api/v1/products
	router.PartyFunc("/products", func(router iris.Party) {		
		router.Get("/", deps.Handler(Products))
		//router.PartyFunc("/{name:string}", registerProductRoutes)
	})

}
