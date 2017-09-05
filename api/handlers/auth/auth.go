package auth

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mikunalpha/mog/api/handlers"
	"github.com/mikunalpha/mog/api/router/middlewares"
	"github.com/mikunalpha/mog/api/shared/auth"
	"github.com/mikunalpha/mog/api/shared/errors"
	"github.com/mikunalpha/mog/api/shared/store"
)

// Register binds handlers to routing rules.
func Register(rg *gin.RouterGroup) {
	v1Router := rg.Group("/v1", middlewares.Authenticate(), middlewares.Services())
	v1Router.GET("/auth/info", GetAuthInfo)
	v1Router.POST("/auth/login", PostLogin)
}

// GetAuthInfo responses a account data according to token.
func GetAuthInfo(c *gin.Context) {
	authInfo := c.MustGet("authInfo").(*auth.AuthInfo)

	// Populate basic authInfo.
	ai := auth.AuthInfo{
		Id:   authInfo.Id,
		Role: authInfo.Role,
	}

	// Response authInfo.
	c.JSON(200, struct {
		Data auth.AuthInfo `json:"data"`
	}{
		ai,
	})
}

// PostLogin set jwt token cookie of a valid account.
func PostLogin(c *gin.Context) {
	var err error
	var req CredentialsData

	// Parse JSON body into struct.
	err = c.BindJSON(&req)
	if err != nil {
		handlers.Abort(c, errors.ParseError.SetOriginError(err))
		// delete cookie access token
		c.SetCookie("at", "", 0, "/", "", false, false)
		return
	}
	if req.Data == nil {
		handlers.Abort(c, errors.ParamError)
		// delete cookie access token
		c.SetCookie("at", "", 0, "/", "", false, false)
		return
	}

	account, err := c.MustGet("store").(store.Store).
		Accounts().FindByCredentials(req.Data.Username, req.Data.Password)
	if err != nil {
		handlers.Abort(c, errors.NotFoundError.SetOriginError(err).ReplaceMessage("Provided username or password was incorrect."))
		// delete cookie access token
		c.SetCookie("at", "", 0, "/", "", false, false)
		return
	}

	// Generate token.
	ai := auth.AuthInfo{}
	ai.Id = *account.Id
	ai.Role = *account.Role
	ai.ExpiresAt = time.Now().Add(time.Hour * 24 * 90).Unix()
	encryptedToken, err := auth.EncryptToken(ai)
	if err != nil {
		handlers.Abort(c, errors.ServerError.SetOriginError(err))
		// delete cookie access token
		c.SetCookie("at", "", 0, "/", "", false, false)
		return
	}

	// Set encryptedToken into Cookie 90 days.
	c.SetCookie("at", encryptedToken, 7776000, "/", "", false, false)

	// Response encrypted access token and role.
	c.JSON(200, struct {
		Data interface{} `json:"data"`
	}{
		struct {
			AccessToken string `json:"access_token"`
			Role        int64  `json:"role"`
		}{
			encryptedToken,
			*account.Role,
		},
	})
}
