package services

import (
	"github.com/BurntSushi/toml"
	"gocherry-api-gateway/components/common_enum"
	"log"
	"os"
	"strings"
)

func GetAppConfig() *common_enum.Config {
	dir, _ := os.Getwd()
	dir = strings.Replace(dir, "proxy", "", -1)
	dir = strings.Replace(dir, "admin", "", -1)
	config := &common_enum.Config{}
	_, err := toml.DecodeFile(dir + "/config/app.tml", config)
	if err != nil {
		log.Println("配置文件读取错误")
	}
	appConfig := config
	return appConfig
}
