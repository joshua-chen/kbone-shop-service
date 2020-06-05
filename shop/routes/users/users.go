/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-15 22:46:07
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-28 21:13:43
 */
package users

import (
	"github.com/joshua-chen/go-commons/mvc/context/request"
	"github.com/joshua-chen/go-commons/middleware/models"
	"github.com/joshua-chen/go-commons/mvc/context/response"
	_ "github.com/joshua-chen/go-commons/utils/security/aes"
	_ "errors"
	"shop/services"
	_ "time"

	"github.com/kataras/iris/v12"

)

// @Summary 获取用户信息
// @Description 登录
// @Produce json
// @Success 200 {string} string "ok" "返回用户信息"
// @Failure 400 {string} string "err_code：10002 参数错误； err_code：10003 校验错误"
// @Failure 401 {string} string "err_code：10001 登录失败"
// @Failure 500 {string} string "err_code：20001 服务错误；err_code：20002 接口错误；err_code：20003 无数据错误；err_code：20004 数据库异常；err_code：20005 缓存异常"
// @Router /common/users [get]
func Users(ctx iris.Context, service services.UserService) (result response.Result) {
	page := request.NewPagination(ctx)
	data, total := service.GetAll(page)
	//return response.ToResult(data)
	return response.DefaultResult(iris.Map{"rows": data, "total": total})

}
// @Summary 用户注册
// @Description 用户注册
// @Produce json
// @Success 200 {string} string "ok" "返回用户信息"
// @Failure 400 {string} string "err_code：10002 参数错误； err_code：10003 校验错误"
// @Failure 401 {string} string "err_code：10001 登录失败"
// @Failure 500 {string} string "err_code：20001 服务错误；err_code：20002 接口错误；err_code：20003 无数据错误；err_code：20004 数据库异常；err_code：20005 缓存异常"
// @Router /common/users/registe [post]
func Registe(ctx iris.Context, service services.UserService) (result response.Result) {
	user := new(models.User)
	ctx.ReadJSON(&user)
	return response.DefaultResult(service.Registe(user))
}

func  NewToken(ctx iris.Context, service services.UserService)  (result response.Result){
	username := ctx.Params().Get("username")
	user := models.User{Username:username}
	token:= service.NewToken(&user)
	return response.DefaultResult(token)
}
