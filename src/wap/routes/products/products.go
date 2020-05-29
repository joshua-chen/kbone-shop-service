/*
 * @Descripttion: 
 * @version: 
 * @Author: joshua
 * @Date: 2020-05-18 14:54:08
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-28 21:13:34
 */ 
package products

import (
	"commons/mvc/context/request"
	"commons/mvc/context/response"
	_ "commons/mvc/models"
	_ "shop/datamodels"
	"shop/services"

	"github.com/kataras/iris/v12"

)


// @Summary 产品列表
// @Description 产品列表
// @Produce json
// @Param body body request.Pagination true "body参数"
// @Success 200 {string} string "ok" "返回用户信息"
// @Failure 400 {string} string "err_code：10002 参数错误； err_code：10003 校验错误"
// @Failure 401 {string} string "err_code：10001 登录失败"
// @Failure 500 {string} string "err_code：20001 服务错误；err_code：20002 接口错误；err_code：20003 无数据错误；err_code：20004 数据库异常；err_code：20005 缓存异常"
// @Router /wap/products [post]
func Products(ctx iris.Context, service services.ProductService) (result response.Result) {
	page := request.NewPagination(ctx)
	data, count := service.GetAll(page)
	//return response.ToResult(data)
	return response.DefaultResult(iris.Map{"rows": data, "count": count})

}
