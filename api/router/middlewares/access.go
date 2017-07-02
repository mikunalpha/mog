package middlewares

import (
	"log"
	"net/url"
	"strconv"
	"time"

	"gopkg.in/gin-gonic/gin.v1"
)

type accessLog struct {
	LogType               string  `json:"t"`
	ClientIP              string  `json:"client_ip"`
	Method                string  `json:"method"`
	StatusCode            int     `json:"status_code"`
	Path                  string  `json:"path"`
	Related               string  `json:"related"`
	Query                 string  `json:"query"`
	Latency               float64 `json:"latency"`
	RequestContentLength  uint64  `json:"req_len"`
	ResponseContentLength uint64  `json:"resp_len"`
	Platform              string  `json:"platform"`
	Model                 string  `json:"model"`
	Application           string  `json:"application"`
}

// Access logs api requests.
func Access() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start time
		start := time.Now()

		c.Next()

		latency := time.Now().Sub(start)

		q, _ := url.ParseQuery(c.Request.URL.RawQuery)

		l := accessLog{
			LogType:     "access",
			ClientIP:    c.ClientIP(),
			Method:      c.Request.Method,
			StatusCode:  c.Writer.Status(),
			Latency:     latency.Seconds(),
			Path:        c.Request.URL.Path,
			Related:     q.Get("related"),
			Query:       c.Request.URL.RawQuery,
			Platform:    c.Request.Header.Get("Platform"),
			Model:       c.Request.Header.Get("Model"),
			Application: c.Request.Header.Get("Application"),
		}

		if u, err := strconv.ParseUint(c.Request.Header.Get("Content-Length"), 10, 64); err == nil {
			l.RequestContentLength = u
		}
		l.ResponseContentLength = uint64(c.Writer.Size())

		if c.Request.URL.RawQuery != "" {
			l.Path = l.Path + "?" + c.Request.URL.RawQuery
		}

		if l.StatusCode == 200 {
			log.Printf("(%fs) %s [%s] %s [%d]", l.Latency, l.ClientIP, l.Method, l.Path, l.StatusCode)
		} else if l.StatusCode >= 300 && l.StatusCode < 400 {
			log.Printf("(%fs) %s [%s] %s [%d]", l.Latency, l.ClientIP, l.Method, l.Path, l.StatusCode)
		} else if l.StatusCode >= 400 && l.StatusCode < 500 {
			log.Printf("(%fs) %s [%s] %s [%d]", l.Latency, l.ClientIP, l.Method, l.Path, l.StatusCode)
		} else if l.StatusCode >= 500 {
			log.Printf("(%fs) %s [%s] %s [%d]", l.Latency, l.ClientIP, l.Method, l.Path, l.StatusCode)
		} else {
			log.Printf("(%fs) %s [%s] %s [%d]", l.Latency, l.ClientIP, l.Method, l.Path, l.StatusCode)
		}
	}
}
