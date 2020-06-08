package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		var params string
		if c.Request.Method == "GET" {
			params = c.Request.URL.RawQuery
		} else {
			params = c.Request.PostForm.Encode()
		}
		log.Printf(fmt.Sprintf("%v -- %v -- %v -- %s -- %s",
			c.Writer.Status(),
			latency,
			c.Request.Method,
			c.Request.URL.Path,
			params,
		))
	}
}
