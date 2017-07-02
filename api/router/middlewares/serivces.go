package middlewares

import (
	"github.com/mikunalpha/mog/api/services/store"
	"gopkg.in/gin-gonic/gin.v1"
)

// Services assigns instances of services for handlers.
func Services() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Store service
		s := store.S.Copy()

		c.Set("store", s)

		c.Next()

		s.Close()
	}
}
