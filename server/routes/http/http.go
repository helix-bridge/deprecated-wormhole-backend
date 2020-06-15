package http

import (
	"github.com/gin-gonic/gin"
)

func Run(api *gin.RouterGroup) {
	api.GET("supply", supply())
	api.GET("ringBurn", ringBurn())
}

func JsonFormat(data interface{}, code int) map[string]interface{} {
	r := gin.H{
		"data": data,
		"code": code,
		"msg":  responseCode[code],
	}
	return r
}

var responseCode = map[int]string{
	0:    "ok",
	1001: "params error",
}