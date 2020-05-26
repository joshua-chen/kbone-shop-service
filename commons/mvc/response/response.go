/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-25 17:37:30
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-26 18:36:10
 */

package response

import (
	"commons/mvc/models"

)

func ToResult(data interface{}) *models.ResponseResult {

	//var result = new(models.ResponseResult)
	//result := models.NewResponseResult(data, "200")
	return models.NewResponseResult(data, "200")
	//return result
}
