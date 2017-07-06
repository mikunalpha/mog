package middlewares

import (
	"log"
	"strings"

	"github.com/mikunalpha/mog/api/shared/auth"
	"github.com/mikunalpha/mog/api/shared/errors"
	"gopkg.in/gin-gonic/gin.v1"
)

// Authenticate decrypts token and set authInfo.
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		var encryptedToken string

		// Get token from cookie.
		encryptedTokenFromCookie, err := c.Cookie("at")

		// Get token from header.
		s := strings.Split(c.Request.Header.Get("Authorization"), " ")
		if len(s) == 2 && s[0] == "Bearer" {
			encryptedToken = s[1]
		}

		if err == nil && encryptedToken != "" {
			// Check whether token from header and cookie are same.
			if encryptedTokenFromCookie != encryptedToken {
				log.Println("Token from header and cookie are not same.")
				c.JSON(
					errors.AuthenticationError.Status,
					errors.AuthenticationError.ReplaceMessage("Cannot validate your token."))
				c.Abort()
				return
			}

			// Get claims in the token.
			authInfo, err := auth.DecryptToken(encryptedToken)
			if err != nil {
				log.Println("Error occured when decrypting token.", err)
				c.JSON(
					errors.AuthenticationError.Status,
					errors.AuthenticationError.ReplaceMessage("Cannot identify your token."))
				c.Abort()
				return
			}

			if authInfo.Id == "" || authInfo.Role < auth.Anonymous {
				c.JSON(
					errors.AuthenticationError.Status,
					errors.AuthenticationError.ReplaceMessage("Cannot identify invalid token."))
				c.Abort()
				return
			}

			// Set authInfo.
			c.Set("authInfo", authInfo)
		} else {
			// Set authInfo as anonymous.
			c.Set("authInfo", &auth.AuthInfo{
				Role: auth.Anonymous,
			})
		}

		c.Next()
	}
}
