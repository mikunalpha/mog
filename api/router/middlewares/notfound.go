package middlewares

import (
	"log"

	"github.com/mikunalpha/mog/api/shared/errors"
	"gopkg.in/gin-gonic/gin.v1"
)

// Notfound deal with not found requests.
func Notfound() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("An error occurred because the routing rule is not found.")

		c.JSON(
			errors.NotFoundError.Status,
			errors.NotFoundError.ReplaceMessage("An error occurred because the routing rule is not found."))
		c.Abort()
	}
}
