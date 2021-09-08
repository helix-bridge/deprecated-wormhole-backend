package http

import (
	"github.com/darwinia-network/link/db"
	"github.com/darwinia-network/link/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ringBurn() gin.HandlerFunc {
	return func(c *gin.Context) {
		p := new(struct {
			Address string `json:"address" binding:"required" form:"address"`
			Page    int    `json:"page" form:"page"`
			Row     int    `json:"row" form:"row"`
		})
		if err := c.ShouldBindQuery(p); err != nil {
			c.JSON(http.StatusOK, JsonFormat(nil, 1001))
			return
		}
		if !util.VerifyTronAddress(p.Address) && !util.VerifyEthAddress(p.Address) {
			c.JSON(http.StatusOK, JsonFormat(nil, 1001))
			return
		}
		list, count := db.RingBurnList(p.Address, p.Page, p.Row)
		c.JSON(http.StatusOK, JsonFormat(map[string]interface{}{
		    "list": list, "count": count,
		}, 0))
	}
}
