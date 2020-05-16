package config

import (
	"encoding/json"
	"os"
)

type AppConfig struct {
	AppName    string `json:"appName"`
	Port       string `json:"port"`
	StaticPath string `json:"staticPath"`
	Env        string `json:"env"`
}

var ServConfig AppConfig

func GetAppConfig() *AppConfig {
	file, err := os.Open("config/app.json")
	if err != nil {
		panic(err.Error())
	}
	decoder := json.NewDecoder(file)
	conf := AppConfig{}
	err = decoder.Decode(&conf)
	if err != nil {
		panic(err.Error())
	}
	return &conf

}
