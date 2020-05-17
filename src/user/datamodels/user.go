/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-17 00:11:14
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-17 19:14:48
 */
package datamodels

import (
	commons "commons/datamodels"
)

type User struct {
	commons.User
	Name   string //姓名
	Email  string //邮箱
	Mobile string //手机
	QQ     string
	Gender int    //0男 1女
	Age    int    //年龄
	Remark string //备注

}
