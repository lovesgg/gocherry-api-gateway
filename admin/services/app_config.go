package services

import (
	"github.com/BurntSushi/toml"
	"gocherry-api-gateway/components/common_enum"
	"log"
	"os"
	"strings"
)

func GetAppConfig() *common_enum.Config {
	dir,_ := os.Getwd()
	dir = strings.Replace(dir, "admin", "", -1)
	dir = strings.Replace(dir, "proxy", "", -1)
	config := &common_enum.Config{}
	path := dir + "/config/app.tml"
	_, err := toml.DecodeFile(path, config)
	if err != nil {
		log.Println("配置文件读取错误")
	}
	appConfig := config
	return appConfig
}
