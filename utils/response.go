package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
}

func ResponseSuccess(c *gin.Context, data any) {
	c.JSON(200, response{
		Code: 200,
		Msg:  "success",
		Data: data,
	})
}

func ResponseError(c *gin.Context, msg string) {
	c.JSON(200, response{
		Code: 500,
		Msg:  msg,
	})
}

func ResponseBadRequest(c *gin.Context, msg string) {
	c.JSON(400, response{
		Code: 400,
		Msg:  msg,
	})
}

func ResponseErrorWithHttpCode(c *gin.Context, code int, msg string) {
	c.JSON(code, response{
		Code: code,
		Msg:  msg,
	})
}

func DefaultQueryInt(c *gin.Context, key string, defaultValue int) int {
	value, _ := strconv.Atoi(c.DefaultQuery(key, strconv.Itoa(defaultValue)))
	return value
}

func DefaultQueryString(c *gin.Context, key string, defaultValue string) string {
	return c.DefaultQuery(key, defaultValue)
}
