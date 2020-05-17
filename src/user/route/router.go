/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-17 00:11:13
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-17 16:45:16
 */
/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-16 23:39:39
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-17 00:14:45
 */

package route

import (
	"../controllers"
	"commons/mvc/models"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"net/http"
)

func InitRouter(app *iris.Application) {
	//app.Use(CrossAccess)
	baseUri := "/api/user/v1"
	mvc.New(app.Party(baseUri + "/users")).Handle(controllers.NewUserController())

}

func CrossAccess11(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}
func CrossAccess(ctx iris.Context) {
	ctx.ResponseWriter().Header().Add("Access-Control-Allow-Origin", "*")
}
