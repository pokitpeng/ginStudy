package common

import (
	"fmt"
	"ginStudy/model"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// 连接数据库
func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "127.0.0.1"
	port := "3306"
	database := "gin"
	username := "root"
	password := "123456"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", username, password, host, port, database, charset)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("faild to connect database,err:" + err.Error())
	}
	// 如果数据表不存在，自动创建
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
