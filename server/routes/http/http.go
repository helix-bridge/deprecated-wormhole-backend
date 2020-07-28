package http

import (
	"github.com/gin-gonic/gin"
)

func Run(server *gin.Engine) {
	server.GET("supply", ringSupply())
	server.GET("supply/kton", ktonSupply())
	api := server.Group("/api")
	api.GET("ringBurn", ringBurn())
	api.GET("/status", func(c *gin.Context) {
		c.String(200, "OK")
	})
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
