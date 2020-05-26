/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-15 22:46:07
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-25 18:46:11
 */
package routes

import (
	"shop/datamodels"
	"shop/services"

)

// @description 获取用户
// @version 1.0
// @accept application/x-json-stream
// @router /users [get]
func Users(service services.UserService) (results []datamodels.User) {
	return service.GetAll()
}
