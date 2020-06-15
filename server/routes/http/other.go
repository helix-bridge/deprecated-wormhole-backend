package http

import (
	"github.com/darwinia-network/link/db"
	"github.com/darwinia-network/link/services/parallel"
	"github.com/darwinia-network/link/util"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"math/big"
	"net/http"
)

func supply() gin.HandlerFunc {
	return func(c *gin.Context) {
		supply := new(big.Int).Add(parallel.RingEthSupply(), parallel.RingTronSupply())
		c.JSON(http.StatusOK,
			JsonFormat(decimal.NewFromBigInt(supply, 0).Div(decimal.New(1, 18)), 0),
		)
	}
}

func ringBurn() gin.HandlerFunc {
	return func(c *gin.Context) {
		p := new(struct {
			Address string `json:"address" binding:"required" form:"address"`
		})
		if err := c.ShouldBindQuery(p); err != nil {
			c.JSON(http.StatusOK, JsonFormat(nil, 1001))
			return
		}
		if !util.VerifyTronAddress(p.Address) && !util.VerifyEthAddress(p.Address) {
			c.JSON(http.StatusOK, JsonFormat(nil, 1001))
			return
		}
		c.JSON(http.StatusOK, JsonFormat(db.RingBurnList(p.Address), 0))
	}
}
