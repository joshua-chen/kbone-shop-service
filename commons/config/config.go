/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-18 16:46:19
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-22 14:13:28
 */
/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-18 16:46:19
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-22 11:47:05
 */

package config

import (
	_ "bytes"
	"commons/utils/yaml"
	_ "compress/gzip"
	_ "fmt"
	_ "io"
	_ "io/ioutil"
	_ "os"
	_ "path/filepath"
	_ "strings"
	_ "time"

	"github.com/kataras/golog"
	_ "gopkg.in/yaml.v2"
)

func init() {
	//GetAppConfig()
	//GetDbConfig()
	yaml.ReadYaml("config/app.yml", &AppConfig)
	yaml.ReadYaml("config/db.yml", &DBConfig)
	//initRootUser()
	golog.Info("[DBConfig]==> ", DBConfig)
	golog.Info("[AppConfig]==> ", AppConfig)

}
