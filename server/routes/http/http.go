package http

import (
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"time"
)

func Run(server *gin.Engine) {

	store := persistence.NewInMemoryStore(time.Second)

	server.GET("supply/ring", cache.CachePage(store, time.Minute, ringSupply()))
	server.GET("supply", cache.CachePage(store, time.Minute, ringSupply()))
	server.GET("supply/kton", cache.CachePage(store, time.Minute, ktonSupply()))
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
