package middlewares

import (
	"log"
	"runtime"

	"github.com/mikunalpha/mog/api/shared/errors"
	"gopkg.in/gin-gonic/gin.v1"
)

// Recovery deals with painc.
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)

				for i := 3; ; i++ {
					_, file, line, ok := runtime.Caller(i)
					if !ok {
						break
					}
					log.Printf("%s:%d", file, line)
				}
				c.JSON(
					errors.ServerError.Status,
					errors.ServerError)
				c.Abort()
			}
		}()

		c.Next()
	}
}
