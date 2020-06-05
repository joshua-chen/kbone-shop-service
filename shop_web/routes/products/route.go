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
	_"github.com/joshua-chen/go-commons/middleware"
	_"github.com/joshua-chen/go-commons/middleware/jwt"
	 "github.com/joshua-chen/go-commons/mvc/route"
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

 
func Register(app *iris.Application) {
	
	//api.PartyFunc("/{name:string}", registerProductRoutes)
	deps := hero.New()
	service := services.NewProductService()
	deps.Register(service)
	
	// GET: /api/v1/products
	 
	route.PartyWeb(app, func(router iris.Party) {
		router.PartyFunc("/products", func(router iris.Party){
			router.Get("/", deps.Handler(Products))
			//router.PartyFunc("/{name:string}", registerProductRoutes)
		})	 
	})
}
 