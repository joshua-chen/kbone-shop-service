/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-16 23:39:39
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-28 21:30:16
 */

package users

import (
	_"github.com/joshua-chen/go-commons/middleware"
	_"github.com/joshua-chen/go-commons/middleware/jwt"
	_ "github.com/joshua-chen/go-commons/mvc/route"
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
	service := services.NewUserService()
	deps.Register(service)
	
	// GET: /api/v1/common/users
	route.PartyCommon(app, func(router iris.Party) {
		router.PartyFunc("/users", func(router iris.Party){
			router.Get("/", deps.Handler(Users))
			router.Any("/registe", deps.Handler(Registe))
			router.Any("/newtoken", deps.Handler(NewToken))

			//NewToken
			//router.PartyFunc("/{name:string}", registerProductRoutes)
		})	 
	})
}
 
