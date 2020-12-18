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
	server.GET("supply/ring", cache.CachePage(store, time.Minute*5, ringSupply()))
	server.GET("supply/kton", cache.CachePage(store, time.Minute*5, ktonSupply()))
	api.GET("/status", func(c *gin.Context) {
		c.String(200, "OK")
	})
	api.GET("ringBurn", ringBurn())
	api.GET("redeem", redeem())
	api.GET("redeem/stat", redeemStat())
	api.GET("supply", cache.CachePage(store, time.Minute, ringSupply()))
	api.GET("ethereumBacking/locks",locks())
	api.POST("/subscribe", subscribe())
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
