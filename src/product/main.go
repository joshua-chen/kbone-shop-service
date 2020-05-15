package main

import (
	"commons/config"
	_ "time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	_ "github.com/kataras/iris/mvc"
	_ "github.com/kataras/iris/sessions"

)

func main() {
	app := newApp()

	//应用App设置
	configation(app)

	//路由设置
	mvcHandle(app)
	config := config.GetConfig()
	//addr := "localhost:" + config.Port
	app.Run(
		iris.Addr(":"+config.Port),                    //在端口8080进行监听
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

//MVC 架构模式处理
func mvcHandle(app *iris.Application) {
	//启用session
	/*sessionManager := sessions.New(sessions.Config{
		Cookie:  "sessioncookie",
		Expires: 24 * time.Hour,
	})
	
		engine := datasource.NewMysqlEngine()


		//管理员模块功能
		adminService := service.NewAdminService(engine)

		admin := mvc.New(app.Party("/admin")) //设置路由组
		admin.Register(
			adminService,
			sessionManager.Start,
		)
		//通过mvc的Handle方法进行控制器的指定
		admin.Handle(new(controller.AdminController))
	*/
}
