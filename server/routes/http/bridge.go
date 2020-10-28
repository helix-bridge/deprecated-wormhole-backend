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

func EthereumRelay() gin.HandlerFunc {
	return func(c *gin.Context) {
		p := new(struct {
			Tx         string `json:"tx" binding:"required"`
			DarwiniaTx string `json:"darwinia_tx" binding:"required"`
		})
		if err := c.ShouldBindJSON(p); err != nil {
			c.JSON(http.StatusOK, JsonFormat(nil, 1001))
			return
		}
		db.UpdateRedeem(p.Tx,p.DarwiniaTx)
		c.JSON(http.StatusOK,JsonFormat(nil, 0))
	}
}
