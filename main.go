package main

import (
	"ginStudy/common"
	_ "ginStudy/docs"
	"ginStudy/logger"
	"ginStudy/router"
	"ginStudy/util"
	"log"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	// _ "github.com/go-sql-driver/mysql"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
func main() {
	InitConfig()
	logger.InitLogger()

	db := common.InitDB()
	defer db.Close()

	r := gin.Default()
	r = router.CollectRoute(r)
	// use ginSwagger middleware to
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := viper.GetString("server.port")
	if port == "" {
		log.Fatal("端口号为空")
	}
	log.Printf("swagger接口文档地址：http://127.0.0.1:%v/swagger/index.html", port)
	r.Run(":" + port) // listen and serve on 0.0.0.0:8080

}

// InitConfig 读取配置文件
func InitConfig() {
	workDir, _ := util.GetExecPath() // go build 后使用
	// workDir, _ := os.Getwd()  // go run 调试使用
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
