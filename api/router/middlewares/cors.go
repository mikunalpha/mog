package middlewares

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

// Cors implements the Cross-origin resource sharing which
// is a recommended standard of the W3C.
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Deals with the OPTIONS request.
		c.Writer.Header().Add(
			"Access-Control-Allow-Origin",
			"*",
		)
		c.Writer.Header().Add(
			"Access-Control-Allow-Credentials",
			"true",
		)
		c.Writer.Header().Set(
			"Access-Control-Allow-Methods",
			"GET, POST, PATCH, PUT, DELETE, OPTIONS",
		)
		c.Writer.Header().Set(
			"Access-Control-Allow-Headers",
			"Origin,Accept,Content-Type,Authorization",
		)
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
		}
		c.Next()
	}
}
