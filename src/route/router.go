/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-16 23:39:39
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-18 15:05:32
 */

package route

import (
	"commons/middleware/jwt/route"
	"commons/mvc/models"
	"net/http"
	"shop/controllers"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func InitRouter(app *iris.Application) {
	baseUri := "/api/v1/user"
	mvc.New(app.Party(baseUri + "/users")).Handle(controllers.NewUserController())
}
