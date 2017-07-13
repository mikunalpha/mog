package accounts

import (
	"fmt"

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
	v1Router.POST("/account", Post)
	v1Router.PATCH("/account/:id", Patch)
}

// Post accepts request and create an new account according to request.
// Then, it responses origin request with new id.
func Post(c *gin.Context) {
	var err error
	var req AccountData

	// Parse JSON body into struct.
	err = c.BindJSON(&req)
	if err != nil {
		handlers.Abort(c, errors.ParseError.SetOriginError(err))
		return
	}
	if req.Data == nil {
		handlers.Abort(c, errors.ParamError)
		return
	}

	accountStore := c.MustGet("store").(store.Store).Accounts()

	// Check role number of a new account.
	if req.Data.Role != nil {
		switch {
		case *req.Data.Role < auth.Anonymous || *req.Data.Role > auth.Admin:
			handlers.Abort(c, errors.AuthorizationError.SetOriginError(fmt.Errorf("Cannot create an account with role number %d over range.", *req.Data.Role)))
			return
		case *req.Data.Role == auth.Admin:
			n, err := accountStore.Count()
			if err != nil {
				handlers.Abort(c, errors.DatabaseError.SetOriginError(err))
				return
			}
			if n != 0 {
				handlers.Abort(c, errors.AuthorizationError.SetOriginError(fmt.Errorf("Someone would like to create an admin account.")))
				return
			}
		default:
			handlers.Abort(c, errors.AuthorizationError.SetOriginError(fmt.Errorf("Cannot create an account with invalid role number %d.", *req.Data.Role)))
			return
		}
	} else {
		handlers.Abort(c, errors.AuthorizationError.SetOriginError(fmt.Errorf("Cannot create an account without role.")))
		return
	}

	err = accountStore.Create(req.Data)
	if err != nil {
		handlers.Abort(c, errors.DatabaseError.SetOriginError(err))
		return
	}

	// Response an account with new id.
	c.JSON(200, AccountData{
		Data: &store.Account{
			Id:        req.Data.Id,
			CreatedAt: req.Data.CreatedAt,
		},
	})
}

// Patch accepts request and update an account.
func Patch(c *gin.Context) {
	var err error
	var req AccountData

	// Check id param from URL path.
	id := c.Param("id")
	if len(id) != 24 {
		handlers.Abort(c, errors.ParseError.SetOriginError(fmt.Errorf("Format of id is invalid.")))
		return
	}

	// Parse JSON body into struct.
	err = c.BindJSON(&req)
	if err != nil {
		handlers.Abort(c, errors.ParseError.SetOriginError(err))
		return
	}
	if req.Data == nil {
		handlers.Abort(c, errors.ParamError)
		return
	}

	accountStore := c.MustGet("store").(store.Store).Accounts()

	_, err = accountStore.Find(c.Param("id"))
	if err == nil {
		handlers.Abort(c, errors.NotFoundError.SetOriginError(err))
		return
	}

	err = accountStore.Update(id, *req.Data)
	if err == nil {
		handlers.Abort(c, errors.DatabaseError.SetOriginError(err))
		return
	}

	// Response an account with new id.
	c.JSON(200, AccountData{
		Data: &store.Account{
			Id: req.Data.Id,
		},
	})
}
