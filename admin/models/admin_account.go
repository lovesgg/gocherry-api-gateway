package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type AdminAccount struct {
	gorm.Model
	Id       int    `json:"id"`
	UserName string `json:"user_name"`
	Phone    string `json:"phone"`
	Pwd      string `json:"pwd"`
	Level    int    `json:"level"`
	State    int    `json:"state"`
}

func (m *AdminAccount) GetUserBydPwd(phone string, password string, admin AdminAccount) AdminAccount {
	db = GetInstance().GetMysqlDB()
	db.Where("phone = ? and pwd = ?", phone, password).First(&admin)
	return admin
}

func (m *AdminAccount) GetUserList(page int, pageSize int, admins []AdminAccount) []AdminAccount {
	db = GetInstance().GetMysqlDB()
	db.Where("level < ? and state = ?", 3, 1).Find(&admins)
	return admins
}

func (m *AdminAccount) SaveUser(account AdminAccount) bool {
	db = GetInstance().GetMysqlDB()
	db.Create(&account)
	return true
}

func (m *AdminAccount) DelUser(account AdminAccount) bool {
	db = GetInstance().GetMysqlDB()
	db.Model(&account).Where("phone = ?", account.Phone).Update("state", 0)

	return true
}
