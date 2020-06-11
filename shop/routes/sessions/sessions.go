/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-15 22:46:07
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-28 21:13:43
 */
package sessions

import (
	_"github.com/joshua-chen/go-commons/mvc/context/request"
	"github.com/joshua-chen/go-commons/middleware/models"
	"github.com/joshua-chen/go-commons/mvc/context/response"
	_ "github.com/joshua-chen/go-commons/utils/security/aes"
	_ "errors"
	"shop/services"
	_ "time"

	"github.com/kataras/iris/v12"

)



func  NewToken(ctx iris.Context, service services.SessionService)  (result response.Result){
	username := ctx.Params().Get("username")
	user := models.User{Username:username}
	token:= service.NewToken(&user)
	return response.DefaultResult(token)
}
