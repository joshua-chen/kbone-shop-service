/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-17 00:11:14
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-27 16:43:49
 */
package models

type User struct {
	ID       int64 `json:"id"`
	Username string `json:"username"`
	Password string
	Token    string `json:"-"`
	Session  string `json:"-"`
}
