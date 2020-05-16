/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-17 00:11:14
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-17 00:28:09
 */
package datamodels

type User struct {
	ID       string `json:"ID"`
	Username string `json:"Username"`
	Password string
	Name     string //姓名
	Email    string //邮箱
	Mobile   string //手机
	QQ       string
	Gender   int    //0男 1女
	Age      int    //年龄
	Remark   string //备注
	Token    string `json:"-"`
	Session  string `json:"-"`
}
