package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

func RequestHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("requestStartTime", time.Now())
	}
}
