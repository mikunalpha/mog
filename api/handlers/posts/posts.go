package posts

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/mikunalpha/mog/api/handlers"
	"github.com/mikunalpha/mog/api/router/middlewares"
	"github.com/mikunalpha/mog/api/shared/auth"
	"github.com/mikunalpha/mog/api/shared/errors"
	"github.com/mikunalpha/mog/api/shared/store"
	"gopkg.in/gin-gonic/gin.v1"
)

// Register binds handlers to routing rules.
func Register(rg *gin.RouterGroup) {
	v1Router := rg.Group("/v1", middlewares.Authenticate(), middlewares.Services())
	v1Router.GET("/posts", Get)
	v1Router.GET("/post/:id", Find)
	v1Router.POST("/post", Post)
	v1Router.PATCH("/post/:id", Patch)
	v1Router.DELETE("/post/:id", Delete)
}

// Get responses posts data.
func Get(c *gin.Context) {
	var err error
	var opts store.QueryOptions
	var posts []*store.Post

	authInfo := c.MustGet("authInfo").(*auth.AuthInfo)

	postsStore := c.MustGet("store").(store.Store).Posts()

	offset, err := strconv.Atoi(c.Query("offset"))
	if err == nil && offset > 0 {
		opts.Offset = &offset
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err == nil && limit > 0 {
		opts.Limit = &limit
	}

	if authInfo.Role == auth.Admin {
		if c.Query("keyword") != "" {
			posts, err = postsStore.GetByFuzzyTitleOrFuzzyContent(c.Query("keyword"), &opts)
		} else {
			posts, err = postsStore.Get(&opts)
		}
	} else {
		if c.Query("keyword") != "" {
			posts, err = postsStore.GetPublishedByFuzzyTitleOrFuzzyContent(c.Query("keyword"), &opts)
		} else {
			posts, err = postsStore.GetPublished(&opts)
		}
	}
	if err != nil {
		handlers.Abort(c, errors.DatabaseError.SetOriginError(err))
		return
	}

	maxExcerptContentTextLength := 100
	if c.Query("excerpt") == "1" {
		for i := range posts {
			document, err := goquery.NewDocumentFromReader(strings.NewReader("<div>" + *posts[i].Content + "</div>"))
			if err != nil {
				continue
			}
			excerptContent := ""
			excerptContentTextLength := 0
			ignore := false
			document.Find("div").First().Children().Each(func(n int, s *goquery.Selection) {
				if ignore {
					return
				}
				if excerptContentTextLength >= maxExcerptContentTextLength {
					excerptContent += "<p>...</p>"
					ignore = true
					return
				}

				outerHtml, err := goquery.OuterHtml(s)
				if err == nil {
					excerptContent += outerHtml
				} else {
					html, err := s.Html()
					if err == nil {
						excerptContent += html
					}
				}
				excerptContentTextLength += len(s.Text())
			})
			*posts[i].Content = excerptContent
		}
	}

	c.JSON(200, PostsData{
		Data: posts,
	})
}

// Find responses a post data.
func Find(c *gin.Context) {
	var err error
	var post *store.Post

	authInfo := c.MustGet("authInfo").(*auth.AuthInfo)

	// Check id param from URL path.
	id := c.Param("id")
	if len(id) != 24 {
		handlers.Abort(c, errors.ParseError.SetOriginError(fmt.Errorf("Format of id is invalid.")))
		return
	}

	postsStore := c.MustGet("store").(store.Store).Posts()

	if authInfo.Role == auth.Admin {
		post, err = postsStore.Find(id)
	} else {
		post, err = postsStore.FindPublished(id)
	}
	if err != nil {
		handlers.Abort(c, errors.NotFoundError.SetOriginError(err))
		return
	}

	c.JSON(200, PostData{
		Data: post,
	})
}

// Post uses request data to create a post and responses
// request data with new id or error result.
func Post(c *gin.Context) {
	var err error
	var req PostData

	authInfo := c.MustGet("authInfo").(*auth.AuthInfo)

	if authInfo.Role != auth.Admin {
		fmt.Println(authInfo.Role)
		handlers.Abort(c, errors.AuthorizationError.SetOriginError(err))
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

	postsStore := c.MustGet("store").(store.Store).Posts()

	err = postsStore.CreateByAuthorId(authInfo.Id, req.Data)
	if err != nil {
		handlers.Abort(c, errors.DatabaseError.SetOriginError(err))
		return
	}

	// Reponse a post with new id.
	c.JSON(200, PostData{
		Data: &store.Post{
			Id:        req.Data.Id,
			CreatedAt: req.Data.CreatedAt,
		},
	})
}

// Patch accepts request and update an post.
func Patch(c *gin.Context) {
	var err error
	var req PostData

	authInfo := c.MustGet("authInfo").(*auth.AuthInfo)

	if authInfo.Role != auth.Admin {
		handlers.Abort(c, errors.AuthorizationError.SetOriginError(err))
		return
	}

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

	postsStore := c.MustGet("store").(store.Store).Posts()

	_, err = postsStore.Find(id)
	if err != nil {
		handlers.Abort(c, errors.NotFoundError.SetOriginError(err))
		return
	}

	err = postsStore.UpdateByAuthorId(authInfo.Id, id, *req.Data)
	if err != nil {
		handlers.Abort(c, errors.DatabaseError.SetOriginError(err))
		return
	}

	// Reponse a post with new id.
	c.JSON(200, PostData{
		Data: &store.Post{
			Id: &id,
		},
	})
}

// Delete removes a post by specific id
func Delete(c *gin.Context) {
	var err error

	authInfo := c.MustGet("authInfo").(*auth.AuthInfo)

	if authInfo.Role != auth.Admin {
		handlers.Abort(c, errors.AuthorizationError.SetOriginError(err))
		return
	}

	// Check id param from URL path.
	id := c.Param("id")
	if len(id) != 24 {
		handlers.Abort(c, errors.ParseError.SetOriginError(fmt.Errorf("Format of id is invalid.")))
		return
	}

	postsStore := c.MustGet("store").(store.Store).Posts()

	_, err = postsStore.Find(id)
	if err != nil {
		handlers.Abort(c, errors.NotFoundError.SetOriginError(err))
		return
	}

	if authInfo.Role == auth.Admin {
		err = postsStore.Delete(id)
	} else {
		err = postsStore.DeleteByAuthorId(authInfo.Id, id)
	}
	if err != nil {
		handlers.Abort(c, errors.DatabaseError.SetOriginError(err))
		return
	}

	// Reponse a post with new id.
	c.JSON(200, PostData{
		Data: &store.Post{
			Id: &id,
		},
	})
}
