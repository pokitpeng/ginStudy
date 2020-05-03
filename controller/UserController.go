package controller

import (
	"ginStudy/common"
	"ginStudy/model"
	"ginStudy/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Register(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	// 参数验证
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
		return
	}
	// 判断手机号是否存在
	if isTelephoneExist(DB, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "该手机号已注册"})
		return
	}
	// 如果名字没有传，给一个随机的字符串并创建用户
	if len(name) == 0 {
		name = util.RandString(10)
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	log.Printf("name:%v\ntelephone:%v\npassword:%v\n", name, telephone, password)

	DB.Create(&newUser)
	// 返回结果
	ctx.JSON(200, gin.H{
		"message": "注册成功",
	})
}

// 判断手机号是否已在数据库中
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
