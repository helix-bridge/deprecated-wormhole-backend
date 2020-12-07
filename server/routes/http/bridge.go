package http

import (
	"github.com/darwinia-network/link/db"
	"github.com/darwinia-network/link/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func redeem() gin.HandlerFunc {
	return func(c *gin.Context) {
		p := new(struct {
			Address string `json:"address" binding:"required" form:"address"`
		})
		if err := c.ShouldBindQuery(p); err != nil {
			c.JSON(http.StatusOK, JsonFormat(nil, 1001))
			return
		}
		if !util.VerifyEthAddress(p.Address) {
			c.JSON(http.StatusOK, JsonFormat(nil, 1001))
			return
		}
		c.JSON(http.StatusOK, JsonFormat(db.RedeemList(p.Address), 0))
	}
}


func redeemStat() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, JsonFormat(db.RedeemStat(), 0))
	}
}
