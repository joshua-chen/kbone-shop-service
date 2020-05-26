/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-16 23:39:39
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-26 18:12:35
 */

package routes

import (
	_ "commons/middleware/jwt/route"
	_"commons/mvc/models"
	_ "net/http"
"shop/repositories"
"shop/services"

"github.com/kataras/iris/v12/hero"
	"github.com/kataras/iris/v12"
	_"github.com/kataras/iris/v12/mvc"

)

func InitRouter(app *iris.Application) {
	baseUri := "/api/v1"
	//mvc.New(app.Party(baseUri + "/users")).Handle(controllers.NewUserController())
	repo := repositories.NewProductRepository()
	// Create our movie service, we will bind it to the movie app's dependencies.
	productService := services.NewProductService(repo)
	hero.Register(productService)

	// Serve our routes with hero handlers.
	app.PartyFunc(baseUri+"/product", func(r iris.Party) {
		r.Get("/", hero.Handler(Products))
		//r.Get("/{name}", hero.Handler(routes.HelloName))
	})
}
