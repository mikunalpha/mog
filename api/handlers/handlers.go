package handlers

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
	"strings"

	"github.com/mikunalpha/mog/api/shared/errors"
	"gopkg.in/gin-gonic/gin.v1"
)

// Abort returns error json response and abort process of gin
func Abort(c *gin.Context, err error) {
	e := reflect.Indirect(reflect.ValueOf(err))

	errorMessage := ""
	if c.Param("ssn") != "" {
		errorMessage += c.Param("ssn") + "  "
	}
	if e.FieldByName("Origin").IsValid() && !e.FieldByName("Origin").IsNil() {
		errorMessage =
			fmt.Sprintf(
				"ClientIP:%s %s (%s) (%s)",
				c.ClientIP(), err.Error(), e.FieldByName("Origin"), Here(2))
	} else {
		errorMessage =
			fmt.Sprintf(
				"ClientIP:%s %s (%s)",
				c.ClientIP(), err.Error(), Here(2))
	}
	log.Println(errorMessage)

	if e.FieldByName("Status").IsValid() {
		c.JSON(int(e.FieldByName("Status").Int()), err)
	} else {
		c.JSON(500, errors.ServerError.ReplaceMessage(err.Error()))
	}
	c.Abort()
}

// Here get description of the location of current process in runtime.
// param  skips[0]  1 is the current caller, 2 is parent caller
func Here(skips ...int) string {
	if len(skips) == 0 {
		skips = append(skips, 1)
	}
	fn, file, line, _ := runtime.Caller(skips[0])
	i := strings.LastIndex(file, "/")
	if i != -1 {
		file = file[i+1:]
	}
	return fmt.Sprintf("In file %s at line %d of function %s", file, line, runtime.FuncForPC(fn).Name())
}
