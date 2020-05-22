/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-15 22:46:07
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-18 15:33:48
 */
package controllers

import (
	"shop/datamodels"
	"shop/services"

)

type UserController struct {
	// Our ProductService, it's an interface which
	// is binded from the main application.
	Service services.UserService
}

// @description 获取产品
// @version 1.0
// @accept application/x-json-stream
// @router /users [get]
func (c *UserController) Get() (results []datamodels.User) {
	return c.Service.GetAll()
}
