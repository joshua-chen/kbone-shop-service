/*
 * @Descripttion: 
 * @version: 
 * @Author: joshua
 * @Date: 2020-05-22 15:36:44
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-26 18:38:00
 */ 
package controller

import (
	"commons/middleware/jwt"
	_ "commons/mvc"
	_ "commons/mvc/models"
	_ "fmt"
	_ "log"
	_ "time"

	_ "github.com/dgrijalva/jwt-go/request"
	_ "github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"

)

func getToken(app *iris.Application) {
	app.Get("/getJWTWithExp", func(ctx iris.Context) {
		tokenString := jwt.NewToken(jwt.User{Username: "xxx"})

		// 返回
		ctx.JSON(tokenString)
	})
}
