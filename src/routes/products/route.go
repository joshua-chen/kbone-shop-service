/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-16 23:39:39
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-28 21:20:15
 */

package products

import (
	_"commons/middleware"
	_"commons/middleware/jwt"
	_ "commons/middleware/jwt/route"
	 "commons/mvc/route"
	_ "fmt"
	_ "net/http"
	_"shop/repositories"
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
 
func Register(app *iris.Application) {
	
	//api.PartyFunc("/{name:string}", registerProductRoutes)
	deps := hero.New()
	service := services.NewUserService()
	deps.Register(service)
	
	// GET: /api/v1/products
	route.ApiParty(app).PartyFunc("/products", func(router iris.Party) {		
		router.Get("/", deps.Handler(Products))
		//router.PartyFunc("/{name:string}", registerProductRoutes)
	})
}
 