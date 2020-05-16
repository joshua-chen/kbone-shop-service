/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-16 23:39:39
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-16 23:49:19
 */

 package route

 import (
	 "../controllers"	
	 "commons/mvc/models"
	 "github.com/kataras/iris"
	 "github.com/kataras/iris/mvc"
	 "net/http"
 )
 
 func InitRouter(app *iris.Application) {
	 //app.Use(CrossAccess)
	 baseUri := "/api/product/v1"
	 mvc.New(app.Party(baseUri + "/products")).Handle(controllers.NewUserController())
	
  }
 
 func CrossAccess11(next http.Handler) http.Handler {
	 return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		 w.Header().Add("Access-Control-Allow-Origin", "*")
		 next.ServeHTTP(w, r)
	 })
 }
 func CrossAccess(ctx iris.Context) {
	 ctx.ResponseWriter().Header().Add("Access-Control-Allow-Origin", "*")
 }
 