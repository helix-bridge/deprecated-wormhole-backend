package http

import (
	"github.com/darwinia-network/link/db"
	"github.com/darwinia-network/link/services/email"
	"github.com/darwinia-network/link/util/address"
	"github.com/gin-gonic/gin"
	"net/http"
)

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

func ploSubscribe() gin.HandlerFunc {
	return func(c *gin.Context) {
		p := new(struct {
			Email   string `form:"email" binding:"required,email" json:"email"`
			Address string `form:"address" binding:"required" json:"address"`
		})
		if err := c.ShouldBind(p); err != nil {
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		if address.Decode(p.Address, 2) == "" {
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		_ = db.CreatePloSubscriber(p.Email, p.Address)
		c.JSON(http.StatusOK, map[string]interface{}{"code": 1, "message": "Success"})
	}
}
