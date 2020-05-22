/*
 * @Descripttion: 
 * @version: 
 * @Author: joshua
 * @Date: 2020-05-19 09:29:27
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-19 09:29:40
 */ 
 func GetJsonConfig() *AppConfig {
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