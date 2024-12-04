package resp

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kackerx/interview/common/code"
)

type response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func HandleSuccess(c *gin.Context, data ...any) {
	var ret any
	if len(data) != 0 {
		ret = data[0]
	}

	c.JSON(http.StatusOK, response{
		Code:    code.ErrSuccess.Code(),
		Message: code.ErrSuccess.Msg(),
		Data:    ret,
	})
}

func HandleErr(c *gin.Context, err error, data ...any) {
	var appErr *code.AppError
	if !errors.As(err, &appErr) {
		appErr = code.ErrServer
	}

	// 某些类型错误记录错误日志
	// if appErr.Equal(code.ErrDBUnknow) {
	log.Println("错误日志: ", appErr.Msg())
	// }

	var ret any
	if len(data) != 0 {
		ret = data[0]
	}

	c.JSON(http.StatusOK, response{
		Code:    appErr.Code(),
		Message: appErr.Msg(),
		Data:    ret,
	})
}
