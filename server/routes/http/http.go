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
	api.GET("supply", cache.CachePage(store, time.Minute, ringSupply()))

	//api.GET("ringBurn", ringBurn())
	//api.GET("redeem", redeem())
	//api.GET("mapping/stat", cache.CachePage(store, time.Minute, mappingStat()))
	//api.GET("supply", cache.CachePage(store, time.Minute, ringSupply()))
	//api.GET("ethereumBacking/locks", locks())
	//api.GET("ethereumBacking/lock", lock())
	//api.GET("ethereumBacking/tokenlock", tokenLock())
	//api.GET("ethereumIssuing/register", erc20RegisterResponse())
	//api.GET("ethereumIssuing/burns", erc20TokenBurns())
	//api.POST("/subscribe", subscribe())
	//api.POST("/plo/subscribe", ploSubscribe())
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
