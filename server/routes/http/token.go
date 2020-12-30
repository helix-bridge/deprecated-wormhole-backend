package http

import (
	"github.com/darwinia-network/link/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ringSupply() gin.HandlerFunc {
	return func(c *gin.Context) {
		supply := db.RingSupply()
		if c.Query("t") == "totalSupply" {
			c.String(http.StatusOK, supply.TotalSupply.String())
			return
		}
		c.JSON(http.StatusOK, JsonFormat(supply, 0))
	}
}
func ktonSupply() gin.HandlerFunc {
	return func(c *gin.Context) {
		supply := db.KtonSupply()
		if c.Query("t") == "totalSupply" {
			c.String(http.StatusOK, supply.TotalSupply.String())
			return
		}
		c.JSON(http.StatusOK, JsonFormat(supply, 0))
	}
}
