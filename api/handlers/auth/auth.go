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

// PostLogin set jwt token cookie of a valid account.
func PostLogin(c *gin.Context) {
	var err error
	var req CredentialsData

	// Parse JSON body into struct.
	err = c.BindJSON(&req)
	if err != nil {
		handlers.Abort(c, errors.ParseError.SetOriginError(err))
		c.SetCookie("at", "", 0, "/", "", false, false)
		return
	}
	if req.Data == nil {
		handlers.Abort(c, errors.ParamError)
		c.SetCookie("at", "", 0, "/", "", false, false)
		return
	}

	accountStore := c.MustGet("store").(store.Store).Accounts()

	account, err := accountStore.FindByCredentials(req.Data.Username, req.Data.Password)
	if err != nil {
		handlers.Abort(c, errors.NotFoundError.SetOriginError(err).ReplaceMessage("Provided username or password was incorrect."))
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
		c.SetCookie("at", "", 0, "/", "", false, false)
		return
	}

	// Set encryptedToken into Cookie 90 days.
	c.SetCookie("at", encryptedToken, 7776000, "/", "", false, false)

	// Response encryptedToken.
	c.JSON(200, struct {
		Data interface{} `json:"data"`
	}{
		struct {
			AccessToken string `json:"access_token"`
			Role        int64  `json:"role"`
		}{encryptedToken, *account.Role},
	})
}

// GetAuthInfo responses a account data according to token.
func GetAuthInfo(c *gin.Context) {
	authInfo := c.MustGet("authInfo").(*auth.AuthInfo)

	if authInfo.Role < auth.Anonymous {
		handlers.Abort(c, errors.AuthenticationError)
		return
	}

	// Populate basic authInfo.
	ai := auth.AuthInfo{
		Id:   authInfo.Id,
		Role: authInfo.Role,
	}

	// Response authInfo.
	c.JSON(200, struct {
		Data auth.AuthInfo `json:"data"`
	}{ai})
}
