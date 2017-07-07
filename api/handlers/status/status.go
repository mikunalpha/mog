package status

import (
	"github.com/mikunalpha/mog/api/handlers"
	"github.com/mikunalpha/mog/api/router/middlewares"
	"github.com/mikunalpha/mog/api/shared/auth"
	"github.com/mikunalpha/mog/api/shared/errors"
	"github.com/mikunalpha/mog/api/shared/store"
	"github.com/mikunalpha/mog/api/shared/utils"
	"gopkg.in/gin-gonic/gin.v1"
)

// Register binds handlers to routing rules.
func Register(rg *gin.RouterGroup) {
	v1Router := rg.Group("/v1", middlewares.Authenticate(), middlewares.Services())
	v1Router.GET("/status", Get)
}

// Get returns status of mog.
func Get(c *gin.Context) {
	var err error

<<<<<<< HEAD
	authInfo := c.MustGet("authInfo").(*auth.AuthInfo)
=======
	authInfo := c.MustGet("authInfo").(auth.AuthInfo)
>>>>>>> da1474513963a90bd1176e353ded31eeb3c42dfc

	accountStore := c.MustGet("store").(store.Store).Accounts()

	adminCount, err := accountStore.AdminCount()
	if err != nil {
		handlers.Abort(c, errors.DatabaseError.SetOriginError(err))
		return
	}

	postsStore := c.MustGet("store").(store.Store).Posts()

	var publishedCount int

	if authInfo.Role == auth.Admin {
		publishedCount, err = postsStore.Count()
	} else {
		publishedCount, err = postsStore.PublishedCount()
	}
	if err != nil {
		handlers.Abort(c, errors.DatabaseError.SetOriginError(err))
		return
	}

	status := Status{}
	if adminCount > 0 {
		status.Administered = utils.NewPtrBool(true)
	} else {
		status.Administered = new(bool)
	}
	status.Blog.Posts.Amount = publishedCount

	c.JSON(200, StatusData{
		Data: &status,
	})
}
