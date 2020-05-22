/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-15 22:46:07
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-16 23:25:38
 */
package controllers

import (
	"shop/datamodels"
	"shop/services"

)

type ProductController struct {
	// Our ProductService, it's an interface which
	// is binded from the main application.
	Service services.ProductService
}

// @description 获取产品
// @version 1.0
// @accept application/x-json-stream
// @router /products [get]
func (c *ProductController) Get() (results []datamodels.Product) {
	return c.Service.GetAll()
}
