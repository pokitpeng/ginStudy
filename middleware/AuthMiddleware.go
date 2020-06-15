package middleware

import (
	"ginStudy/common"
	"ginStudy/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 授权中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取授权头
		tokenString := ctx.GetHeader("Authorization")
		// 验证token格式
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "权限不足"})
			ctx.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, cliams, err := common.ParseToken(tokenString)
		//如果解析失败或者token无效
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "权限不足"})
			ctx.Abort()
			return
		}
		// 验证通过后，获取cliams中的userId
		userID := cliams.UserID
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userID)

		// 用户
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "权限不足"})
			ctx.Abort()
			return
		}
		//用户存在，将user信息写入上下文
		ctx.Set("user", user)
		ctx.Next()
	}
}
