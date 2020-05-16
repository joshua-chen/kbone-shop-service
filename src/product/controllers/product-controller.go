package controllers

import (
	"product/datamodels"
	"product/services"
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
