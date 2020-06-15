package common

import (
	"fmt"
	"ginStudy/model"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// DB 定义一个数据库对象
var DB *gorm.DB

// InitDB 连接数据库
func InitDB() *gorm.DB {
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
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

// GetDB 获取初始化后的数据库对象
func GetDB() *gorm.DB {
	return DB
}
