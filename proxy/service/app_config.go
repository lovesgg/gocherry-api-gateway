package services

import (
	"github.com/BurntSushi/toml"
	"gocherry-api-gateway/components/common_enum"
	"log"
)

func GetAppConfig() *common_enum.Config {
	config := &common_enum.Config{}
	_, err := toml.DecodeFile("../config/app.tml", config)
	if err != nil {
		log.Println("配置文件读取错误")
	}
	appConfig := config
	return appConfig
}
