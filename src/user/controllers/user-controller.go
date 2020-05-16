/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-15 22:46:07
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-17 00:46:42
 */
package controllers

import (
	"user/datamodels"
	"user/services"
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
