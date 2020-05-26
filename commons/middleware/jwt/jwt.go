/*
 * @Descripttion:
 * @version:
 * @Author: sueRimn
 * @Date: 2020-05-16 23:24:17
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-26 18:14:22
 */
package jwt

import (
	"commons/mvc/models"
	"fmt"
	_ "log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	_ "github.com/dgrijalva/jwt-go/request"
	_ "github.com/iris-contrib/middleware/cors"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/spf13/cast"

)

const JwtKey = "kboneshop"

func GetJWT() *jwtmiddleware.Middleware {
	jwtHandler := jwtmiddleware.New(jwtmiddleware.Config{
		//这个方法将验证jwt的token
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			//自己加密的秘钥或者说盐值
			return []byte(JwtKey), nil
		},
		Extractor: jwtmiddleware.FromParameter("token"),
		//加密的方式
		SigningMethod: jwt.SigningMethodHS256,
		//验证未通过错误处理方式

		ErrorHandler: func(ctx context.Context, err error) {
			if err == nil {
				return
			}

			ctx.StopExecution()
			ctx.StatusCode(iris.StatusUnauthorized)
			ctx.JSON(models.ResponseResult{
				Code:    "501",
				Msg:     err.Error(),
				Success: false,
			})
		},
	})
	return jwtHandler
}

//生成token
func NewToken(user User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nick_name": user.Username,                                            //用户信息
		"session":   user.Session,                                             //session
		"id":        user.ID,                                                  //用户信息
		"iss":       "Iris",                                                   //签发者
		"iat":       time.Now().Unix(),                                        //签发时间
		"jti":       "9527",                                                   //jwt的唯一身份标识，主要用来作为一次性token,从而回避重放攻击。
		"exp":       time.Now().Add(10 * time.Hour * time.Duration(1)).Unix(), //过期时间
	})
	tokenString, _ := token.SignedString([]byte(JwtKey))

	return tokenString
}

func ParseToken(tokenString string, key string) (interface{}, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		fmt.Println(err)
		return "", false
	}
}

func GetToken(ctx iris.Context) string {
	token := ctx.GetHeader("Authorization")
	if token != "" && len(token) > 7 {
		token = token[7:]
	}
	return token
}
func GetUserID(token string) int {
	var userId = 0
	if token != "" && token != "undefined" && len(token) > 7 {
		v, _ := ParseToken(token, JwtKey)
		if v != "" {
			userId = cast.ToInt(v.(jwt.MapClaims)["id"])
		}
	}
	return userId
}
