package main

import (
	"ginStudy/common"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	// _ "github.com/go-sql-driver/mysql"
)

func main() {
	InitConfig()

	db := common.InitDB()
	defer db.Close()

	r := gin.Default()
	r = CollectRoute(r)

	port := viper.GetString("server.port")
	if port == "" {
		panic(r.Run())
	}
	r.Run(":" + port) // listen and serve on 0.0.0.0:8080
}

// 读取配置文件
func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
