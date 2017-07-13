package router

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mikunalpha/mog/api/handlers/accounts"
	"github.com/mikunalpha/mog/api/handlers/auth"
	"github.com/mikunalpha/mog/api/handlers/posts"
	"github.com/mikunalpha/mog/api/handlers/status"
	"github.com/mikunalpha/mog/api/router/middlewares"
	"github.com/mikunalpha/mog/api/shared/utils"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableBindValidation()
}

var _router *gin.Engine

// _registers contains funcs to register subrouters.
var _registers = []func(*gin.RouterGroup){
	status.Register,
	auth.Register,
	accounts.Register,
	posts.Register,
}

// methodOverride is used to override the method before request go into gin
func methodOverride(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		method := r.Header.Get("X-HTTP-Method-Override")
		if method != "" {
			r.Method = method
		}
		next.ServeHTTP(w, r)
	})
}

// Start begins the routing serivce.
func Start() {
	_router = gin.New()
	_router.NoRoute(middlewares.Notfound())
	_router.NoMethod(middlewares.Notfound())
	_router.Use(middlewares.Recovery())
	_router.Use(middlewares.Access())

	// Deal with CORS.
	_router.Use(middlewares.Cors())
	_router.OPTIONS("/*any", func(c *gin.Context) {
		c.JSON(200, "")
	})

	mogApisRouter := _router.Group("/mogapis")

	mogApisRouter.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"ping": "pong"})
	})

	// Add routing rules.
	for _, r := range _registers {
		r(mogApisRouter)
	}

	// Deal with method override before requests go into gin routing
	handler := methodOverride(_router)
	address := utils.EnvString("ADDRESS", "127.0.0.1:9098")

	stop := make(chan os.Signal)

	// Start daemon.
	s := &http.Server{Addr: address, Handler: handler}
	go func() {
		log.Printf("Start at %s", address)
		if err := s.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	<-stop

	// Gracefully shutdown daemon.
	s.Shutdown(context.Background())
}
