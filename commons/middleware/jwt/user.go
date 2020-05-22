/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-17 00:11:14
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-18 15:44:01
 */
package jwt

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string
	Token    string `json:"-"`
	Session  string `json:"-"`
}
