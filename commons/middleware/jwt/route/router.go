package controller

import (
	"commons/middleware/jwt"
	"commons/mvc"
	"commons/mvc/models"
	_ "fmt"
	_ "log"
	_ "time"

	_ "github.com/dgrijalva/jwt-go/request"
	_ "github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

func getToken(app *mvc.Application) {
	app.Get("/getJWTWithExp", func(ctx iris.Context) {
		tokenString := jwt.NewToken(jwt.User{Username: "xxx"})

		// 返回
		ctx.JSON(tokenString)
	})
}
