/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-15 22:46:07
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-27 13:50:16
 */
package routes

import (
	"commons/mvc/context/request"
	"commons/mvc/context/response"
	_ "commons/mvc/models"
	_ "shop/datamodels"
	"shop/services"

	"github.com/kataras/iris/v12"

)

// @description 获取产品
// @version 1.0
// @accept application/x-json-stream
// @router /products [get]
func Products(ctx iris.Context, service services.ProductService) (result response.Result) {
	page, _ := request.NewPagination(ctx)
	data, count := service.GetAll(page)
	return response.ToResult(iris.Map{"rows": data, "count": count})
}
