/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-18 09:21:47
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-22 14:13:06
 */
package config

import (
	_ "io/ioutil"

	_ "github.com/kataras/golog"
	_ "gopkg.in/yaml.v2"
)

var AppConfig AppInfo

type AppInfo struct {
	Port       string   `yaml:"port"`
	IgnoreURLs []string `yaml:"ignoreUrls"`
	JWTTimeout int64    `yaml:"jwtTimeout"`
	LogLevel   string   `yaml:"logLevel"`
	Secret     string   `yaml:"secret"`
}
