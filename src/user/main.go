package main

import (
	"commons/config"
	"commons/middleware/jwt"
	"github.com/iris-contrib/middleware/cors"
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	_ "github.com/kataras/iris/v12/mvc"
	_ "github.com/kataras/iris/v12/sessions"
	_ "time"
	_ "user/controllers"
	_ "user/docs" //  docs is generated by Swag CLI, you have to import it.

	_ "github.com/betacraft/yaag/irisyaag"
	_ "github.com/betacraft/yaag/yaag"
)

// @title Swagger Product API
// @version 1.0
// @description Product server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url https://github.com/joshua-chen
// @contact.email joshua_chen@qq.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host
// @basePath /api/user/v1
func main() {
	app := newApp()

	//应用App设置
	configation(app)

	appConfig := config.GetAppConfig()
	//addr := "localhost:" + config.Port

	//
	// swaggerUI // use ginSwagger middleware to
	//
	swaggerConfig := &swagger.Config{
		URL: "/swagger/doc.json", //The url pointing to API definition
	}
	// use swagger middleware to
	app.Get("/swagger/{any:path}", swagger.CustomWrapHandler(swaggerConfig, swaggerFiles.Handler))

	/*
		// yaag api 为文档生成器
		yaag.Init(&yaag.Config{
			On:       true,
			DocTitle: "Iris",
			DocPath:  "apidoc.html",
			BaseUrls: map[string]string{"Production": "", "Staging": ""},
		})
		app.Use(irisyaag.New())
		////

	*/
	app.Run(
		iris.Addr(":"+appConfig.Port),                 //在端口8080进行监听
		iris.WithoutServerError(iris.ErrServerClosed), //无服务错误提示
		iris.WithOptimizations,                        //对json数据序列化更快的配置
	)
	//app.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.YAML("./configs/iris.yml")))

}

//构建App
func newApp() *iris.Application {
	app := iris.New()

	//设置日志级别  开发阶段为debug

	app.Logger().SetLevel("debug")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	/*sillyHTTPHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println(r.RequestURI)
	})
	sillyConvertedToIon := iris.FromStd(sillyHTTPHandler)
	app.Use(sillyConvertedToIon)
	*/
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	})
	app.Use(crs)
	app.AllowMethods(iris.MethodOptions)

	app.Use(jwt.GetJWT().Serve) // jwt

	/*
			//注册静态资源
			app.HandleDir("/static", "./static")
			app.HandleDir("/manage/static", "./static")
			app.HandleDir("/img", "./static/img")

			//注册视图文件
			app.RegisterView(iris.HTML("./static", ".html"))
			app.Get("/", func(context context.Context) {
				context.View("index.html")
			})



		authConfig := basicauth.Config{
			Users:   map[string]string{"wangshubo": "wangshubo", "superWang": "superWang"},
			Realm:   "Authorization Required",
			Expires: time.Duration(30) * time.Minute,
		}

		authentication := basicauth.New(authConfig)
	*/

	/*app.Get("/", func(ctx context.Context) { ctx.Redirect("/admin") })

	needAuth := app.Party("/admin", authentication)
	{
		//http://localhost:8080/admin
		needAuth.Get("/", h)
		// http://localhost:8080/admin/profile
		needAuth.Get("/profile", h)

		// http://localhost:8080/admin/settings
		needAuth.Get("/settings", h)
	}
	*/

	return app
}

/**
 * 项目设置
 */
func configation(app *iris.Application) {

	//配置 字符编码
	app.Configure(iris.WithConfiguration(iris.YAML("./config/iris.yml")))

	//错误配置
	//未发现错误
	app.OnErrorCode(iris.StatusNotFound, func(context context.Context) {
		context.JSON(iris.Map{
			"errmsg": iris.StatusNotFound,
			"msg":    " not found ",
			"data":   iris.Map{},
		})
	})

	app.OnErrorCode(iris.StatusInternalServerError, func(context context.Context) {
		context.JSON(iris.Map{
			"errmsg": iris.StatusInternalServerError,
			"msg":    " interal error ",
			"data":   iris.Map{},
		})
	})
}
