package http

import (
	"github.com/darwinia-network/link/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ringSupply() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, JsonFormat(db.RingSupply(), 0))
	}
}
func ktonSupply() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, JsonFormat(db.KtonSupply(), 0))
	}
}
