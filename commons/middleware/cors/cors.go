/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-18 16:03:38
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-18 18:02:46
 */
package cors

import (
	corsmiddleware "github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12/context"

)

func NewCors() context.Handler {
	return corsmiddleware.New(corsmiddleware.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowedMethods:   []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowCredentials: true,
	})
}
