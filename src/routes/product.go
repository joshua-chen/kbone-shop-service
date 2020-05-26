/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-15 22:46:07
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-26 18:35:46
 */
package routes

import (
	_"shop/datamodels"
	"shop/services"
	"commons/mvc/response"
	"commons/mvc/models"

	"github.com/kataras/iris/v12"

)

// @description 获取产品
// @version 1.0
// @accept application/x-json-stream
// @router /products [get]
func Products(ctx iris.Context, service services.ProductService) (result *models.ResponseResult) {
  page ,_:= models.NewPagination(ctx)
	data,_ :=service.GetAll(page)
	return response.ToResult(data)
}
