package http

import (
	"github.com/darwinia-network/link/db"
	"github.com/darwinia-network/link/services/email"
	"github.com/darwinia-network/link/util"
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

func subscribe() gin.HandlerFunc {
	return func(c *gin.Context) {
		p := new(struct {
			Email string `form:"email" binding:"required,email" json:"email"`
		})
		if err := c.ShouldBind(p); err != nil {
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		if err := db.CreateSubscribe(p.Email); err == nil {
			go email.SendToSubscribe(p.Email)
		}
		c.JSON(http.StatusOK, map[string]interface{}{"code": 1, "message": "Success"})
	}
}
