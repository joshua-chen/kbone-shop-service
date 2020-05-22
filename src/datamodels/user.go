/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-17 00:11:14
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-18 15:45:06
 */
package datamodels

import (
	"commons/middleware/jwt"

)

type User struct {
	jwt.User
	Name   string //姓名
	Email  string //邮箱
	Mobile string //手机
	QQ     string
	Gender int    //0男 1女
	Age    int    //年龄
	Remark string //备注

}
