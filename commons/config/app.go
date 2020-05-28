/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-18 09:21:47
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-28 21:55:20
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
	AnonymousUrls []string `yaml:"anonymousUrls"`
	AnonymousRequset AnonymousRequset `yaml:"anonymousRequset"`
	StaticPath  string   `yaml:"staticPath"`
	JwtTimeout int64    `yaml:"jwtTimeout"`
	LogLevel   string   `yaml:"logLevel"`
	Secret     string   `yaml:"secret"`
	ApiPrefix  string   `yaml:"apiPrefix"`
	SqlPath  string   `yaml:"sqlPath"`
}

type AnonymousRequset struct{
	Urls []string `yaml:"urls"`
	Prefixes []string `yaml:"prefixes"`
	Suffixes []string `yaml:"suffixes"`
	
}