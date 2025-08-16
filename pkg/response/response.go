package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	SuccessCode = 0
	ErrorCode   = 1
)

// Success 封装成功的响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: SuccessCode,
		Msg:  "success",
		Data: data,
	})
}

// Fail 封装失败的响应（使用通用错误码）
func Fail(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, Response{
		Code: ErrorCode,
		Msg:  msg,
		Data: nil,
	})
}
