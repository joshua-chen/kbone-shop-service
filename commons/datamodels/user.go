/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-17 00:11:14
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-17 19:15:10
 */
package datamodels

type User struct {
	ID       string `json:"ID"`
	Username string `json:"Username"`
	Password string
	Token    string `json:"-"`
	Session  string `json:"-"`
}
