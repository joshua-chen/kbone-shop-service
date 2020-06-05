/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-16 23:39:39
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-28 21:16:22
 */

package routes

import (
	"shap_wap/routes/products"
	"github.com/kataras/iris/v12"

)



func Register(app *iris.Application) {
	products.Register(app)
	apidoc(app)
}

func apidoc(app *iris.Application) {

	app.Get("/apidoc", func(ctx iris.Context) {
		// 绑定： {{.message}}　为　"Hello world!"
		//ctx.ViewData("message", "Hello world!")
		// 渲染模板文件： ./views/hello.html
		ctx.View("apidoc.html")
	})
}

