package http

import (
	"github.com/darwinia-network/link/config"
	"github.com/darwinia-network/link/db"
	"github.com/darwinia-network/link/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func redeem() gin.HandlerFunc {
	return func(c *gin.Context) {
		p := new(struct {
			Address string `json:"address" binding:"required" form:"address"`
			Page    int    `json:"page" form:"page"`
			Row     int    `json:"row" form:"row"`
			Confirmed string `json:"confirmed" form:"confirmed"`
		})
		if err := c.ShouldBindQuery(p); err != nil {
			c.JSON(http.StatusOK, JsonFormat(nil, 1001))
			return
		}
		if !util.VerifyEthAddress(p.Address) {
			c.JSON(http.StatusOK, JsonFormat(nil, 1001))
			return
		}
		list, count := db.RedeemList(p.Address, p.Page, p.Row, p.Confirmed)
		c.JSON(http.StatusOK, JsonFormat(map[string]interface{}{
		    "list": list, "count": count,
		}, 0))
	}
}

func mappingStat() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, JsonFormat(db.MappingStat(), 0))
	}
}

func locks() gin.HandlerFunc {
	return func(c *gin.Context) {
		p := new(struct {
			Address string `json:"address" binding:"required" form:"address"`
			Page    int    `json:"page" form:"page"`
			Row     int    `json:"row" binding:"required" form:"row"`
			Confirmed string `json:"confirmed" form:"confirmed"`
		})
		if err := c.ShouldBindQuery(p); err != nil {
			c.JSON(http.StatusOK, JsonFormat(nil, 1001))
			return
		}
		if !util.VerifySubstrateAddress(p.Address) {
			c.JSON(http.StatusOK, JsonFormat(nil, 1001))
			return
		}
		list, count := db.DarwiniaBackingLocks(p.Address, p.Page, p.Row, p.Confirmed)
		best, MMRRoot := db.GetMMRIndexBestBlockNum()
		c.JSON(http.StatusOK, JsonFormat(map[string]interface{}{
			"list": list, "count": count, "implName": config.Link.ImplName, "best": best, "MMRRoot": MMRRoot,
		}, 0))
	}
}

func lock() gin.HandlerFunc {
	return func(c *gin.Context) {
		p := new(struct {
			ExtrinsicIndex string `json:"extrinsic_index" binding:"required" form:"extrinsic_index"`
		})
		if err := c.ShouldBindQuery(p); err != nil {
			c.JSON(http.StatusOK, JsonFormat(nil, 1001))
			return
		}
		c.JSON(http.StatusOK, JsonFormat(db.BackingLock(p.ExtrinsicIndex), 0))
	}
}

func erc20RegisterResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		p := new(struct {
			Source string `json:"source" binding:"required" form:"source"`
		})
		if err := c.ShouldBindQuery(p); err != nil {
			c.JSON(http.StatusOK, JsonFormat(nil, 1001))
			return
		}
		if !util.VerifyEthAddress(p.Source) {
			c.JSON(http.StatusOK, JsonFormat(nil, 1002))
			return
		}
		c.JSON(http.StatusOK, JsonFormat(db.TokenRegisterRecordInfo(p.Source), 0))
	}
}

func erc20TokenBurns() gin.HandlerFunc {
	return func(c *gin.Context) {
		p := new(struct {
			Sender string `json:"sender" binding:"required" form:"sender"`
			Page   int    `json:"page" form:"page"`
			Row    int    `json:"row" binding:"required" form:"row"`
			Confirmed string `json:"confirmed" form:"confirmed"`
		})
		if err := c.ShouldBindQuery(p); err != nil {
			c.JSON(http.StatusOK, JsonFormat(nil, 1001))
			return
		}
		if !util.VerifyEthAddress(p.Sender) {
			c.JSON(http.StatusOK, JsonFormat(nil, 1002))
			return
		}
		list, count := db.TokenBurnRecords(p.Sender, p.Page, p.Row, p.Confirmed)
		best, MMRRoot := db.GetMMRIndexBestBlockNum()
		c.JSON(http.StatusOK, JsonFormat(map[string]interface{}{
			"list": list, "count": count, "implName": config.Link.ImplName, "best": best, "MMRRoot": MMRRoot,
		}, 0))
	}
}

func tokenLock() gin.HandlerFunc {
	return func(c *gin.Context) {
		p := new(struct {
			Sender string `json:"sender" binding:"required" form:"sender"`
			Page   int    `json:"page" form:"page"`
			Row    int    `json:"row" binding:"required" form:"row"`
			Confirmed string `json:"confirmed" form:"confirmed"`
		})
		if err := c.ShouldBindQuery(p); err != nil {
			c.JSON(http.StatusOK, JsonFormat(nil, 1001))
			return
		}
		if !util.VerifyEthAddress(p.Sender) {
			c.JSON(http.StatusOK, JsonFormat(nil, 1002))
			return
		}
		list, count := db.EthereumLockList(p.Sender, p.Page, p.Row, p.Confirmed)
		c.JSON(http.StatusOK, JsonFormat(map[string]interface{}{
			"list": list, "count": count, "implName": config.Link.ImplName}, 0))
	}
}

