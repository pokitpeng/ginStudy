package controller

import (
	"ginStudy/common"
	"ginStudy/dto"
	"ginStudy/model"
	"ginStudy/response"
	"ginStudy/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// @Tags 用户模块
// @Summary 注册
// @Description 用户注册接口
// @Accept mpfd
// @Produce json
// @Param name formData string false "用户名"
// @Param telephone formData string true "手机号"
// @Param password formData string true "密码"
// @Success 200 {string} string "{"code":200,"data":...,"message": "注册成功"}"
// @Failure 422 {string} string "{"code":422,"message": ...}"
// @Router /api/v1/auth/register [post]
func Register(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	// 参数验证
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, "手机号必须为11位", gin.H{})
		// ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, "密码不能少于6位", gin.H{})
		// ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
		return
	}
	// 判断手机号是否存在
	if isTelephoneExist(DB, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, "该手机号已注册", gin.H{})
		// ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "该手机号已注册"})
		return
	}
	// 如果名字没有传，给一个随机的字符串并创建用户
	if len(name) == 0 {
		name = util.RandString(10)
	}
	// 加密密码
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, "加密错误", gin.H{})
		// ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "加密错误"})
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	log.Printf("name:%v\ntelephone:%v\npassword:%v\n", name, telephone, password)

	DB.Create(&newUser)
	// 返回结果
	// ctx.JSON(200, gin.H{
	// 	"code":    200,
	// 	"message": "注册成功",
	// })
	response.Response(ctx, http.StatusOK, 200, "注册成功", gin.H{})
}

// @Tags 用户模块
// @Summary 登陆
// @Description 用户登陆接口
// @Accept mpfd
// @Produce json
// @Param telephone formData string true "手机号"
// @Param password formData string true "密码"
// @Success 200 {string} string "{"code":200,"data":...,"message": "登陆成功"}"
// @Failure 422 {string} string "{"code":422,"message": ...}"
// @Router /api/v1/auth/login [post]
func Login(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	// 数据验证
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, "手机号必须为11位", gin.H{})
		// ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, "密码不能少于6位", gin.H{})
		// ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
		return
	}
	// 判断手机号是否存在

	var user model.User
	DB.Where("telephone=?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, "该用户不存在", gin.H{})
		// ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "该用户不存在"})
		return
	}
	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(ctx, http.StatusUnprocessableEntity, 400, "密码错误", gin.H{})
		// ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 400, "msg": "密码错误"})
		return
	}
	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, "系统异常", gin.H{})
		// ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		log.Printf("token generate error:%v", err)
		return
	}
	// 返回结果
	// ctx.JSON(200, gin.H{
	// 	"code":    200,
	// 	"message": "登录成功",
	// 	"data":    gin.H{"token": token},
	// })
	response.Response(ctx, http.StatusOK, 200, "登录成功", gin.H{"token": token})
}

// @Tags 用户模块
// @Summary 获取用户信息
// @Description 获取用户信息
// @Produce json
// @Param name query string false "用户名"
// @Success 200 {string} string "{"code":200,"data":...,"message": "登陆成功"}"
// @Failure 422 {string} string "{"code":422,"message": ...}"
// @Router /api/v1/auth/info [get]
func Info(ctx *gin.Context) {
	//获取的用户肯定是通过认证的，直接先从上下文中获取
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "message": "success", "data": gin.H{"user": dto.ToUserDto(user.(model.User))}})
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
