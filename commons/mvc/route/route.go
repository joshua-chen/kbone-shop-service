/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-22 15:48:01
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-28 21:29:42
 */
package route

import (
	"commons/config"

	"github.com/kataras/iris/v12"

)

func ApiParty(app *iris.Application) iris.Party {
	api := app.Party(config.AppConfig.ApiPrefix)
	return api
}


