package http

import (
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"time"
)

func Run(server *gin.Engine) {

	store := persistence.NewInMemoryStore(time.Second)
	api := server.Group("/api")
	server.GET("supply/ring", cache.CachePage(store, time.Minute, ringSupply()))
	server.GET("supply/kton", cache.CachePage(store, time.Minute, ktonSupply()))
	api.GET("/status", func(c *gin.Context) {
		c.String(200, "OK")
	})
	api.GET("ringBurn", ringBurn())
	api.GET("redeem", redeem())
	api.GET("supply", cache.CachePage(store, time.Minute, ringSupply()))
	api.POST("/subscribe", subscribe())

	internal := server.Group("/internal")
	internal.POST("redeem/relay", EthereumRelay())
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
