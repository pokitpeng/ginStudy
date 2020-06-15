package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 自定义返回
func Response(ctx *gin.Context, httpStatus int, code int, message string, data gin.H) {
	ctx.JSON(httpStatus, gin.H{"code": code, "message": message, "data": data})
}

// Success 常用的成功返回
func Success(ctx *gin.Context, message string, data gin.H) {
	Response(ctx, http.StatusOK, 200, message, data)
}

// Fail 常用的失败返回
func Fail(ctx *gin.Context, message string, data gin.H) {
	Response(ctx, http.StatusOK, 400, message, data)
}
