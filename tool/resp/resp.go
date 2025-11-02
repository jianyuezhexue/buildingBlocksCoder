package resp

import "github.com/gin-gonic/gin"

// 请求成功
func Success(c *gin.Context, data any) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": data,
	})
}

// 请求失败
func Error(c *gin.Context, code int, msg string, data ...any) {
	c.JSON(code, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

// 业务错误
func BizError(c *gin.Context, msg string, data ...any) {
	Error(c, 400, msg, data...)
}
