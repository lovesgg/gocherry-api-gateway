package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gocherry-api-gateway/admin/services"
	"sync"
)

type MysqlConnectiPool struct {
}

var instance *MysqlConnectiPool
var once sync.Once

var db *gorm.DB
var err_db error

func GetInstance() *MysqlConnectiPool {
	once.Do(func() {
		instance = &MysqlConnectiPool{}
	})
	return instance
}

func (m *MysqlConnectiPool) InitDataPool() (issucc bool) {
	appConfig := services.GetAppConfig().Common
	conn := appConfig.MysqlUser + ":" +appConfig.MysqlPass + "@tcp(" + appConfig.MysqlHost + ":" + appConfig.MysqlPort + ")/" + appConfig.MysqlDb

	db, err_db = gorm.Open("mysql", conn)
	db.SingularTable(true)
	db.LogMode(appConfig.MysqlDebug)
	if err_db != nil {
		fmt.Println("err db")
		return false
	}
	return true
}

func (m *MysqlConnectiPool) GetMysqlDB() (db_con *gorm.DB) {
	return db
}
