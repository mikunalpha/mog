package errors

import "fmt"

type Error struct {
	ErrorInfo ErrorInfo `json:"error"`
	Status    int       `json:"-"`
	Origin    error     `json:"-"`
}

type ErrorInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	return e.ErrorInfo.Message
}

// ReplaceMessage changes the error message and return a copy Error
func (e Error) ReplaceMessage(emsg string) Error {
	e.ErrorInfo.Message = emsg
	return e
}

// ReplaceMessagef changes the error message and return a copy Error
func (e Error) ReplaceMessagef(format string, args ...interface{}) Error {
	e.ErrorInfo.Message = fmt.Sprintf(format, args...)
	return e
}

// SetOriginError stores origin error
func (e Error) SetOriginError(oe error) Error {
	e.Origin = oe
	return e
}

func New(i int, c, m string) error {
	return Error{
		ErrorInfo: ErrorInfo{c, m},
		Status:    i,
	}
}

// ParamError
var ParamError = Error{
	ErrorInfo{
		Code:    "ParamError",
		Message: "An error occurred when dealing with param.",
	},
	400,
	nil,
}

// ParseError
var ParseError = Error{
	ErrorInfo{
		Code:    "ParseError",
		Message: "An error occurred when parsing data.",
	},
	400,
	nil,
}

// ValidationError
var ValidationError Error = Error{
	ErrorInfo{
		Code:    "ValidationError",
		Message: "An error occurred when validating data.",
	},
	400,
	nil,
}

// DuplicateError
var DuplicateError Error = Error{
	ErrorInfo{
		Code:    "DuplicateError",
		Message: "An error occurred when operate data.",
	},
	400,
	nil,
}

// AuthenticationError
var AuthenticationError Error = Error{
	ErrorInfo{
		Code:    "AuthenticationError",
		Message: "An error occurred because of unauthenticated operation.",
	},
	403,
	nil,
}

// AuthorizationError
var AuthorizationError Error = Error{
	ErrorInfo{
		Code:    "AuthorizationError",
		Message: "An error occurred because of unauthorized operation.",
	},
	403,
	nil,
}

// NotFoundError
var NotFoundError Error = Error{
	ErrorInfo{
		Code:    "NotFoundError",
		Message: "An error occurred because the data is not found.",
	},
	404,
	nil,
}

// DatabaseError
var DatabaseError Error = Error{
	ErrorInfo{
		Code:    "DatabaseError",
		Message: "An error occurred because of database exception error.",
	},
	500,
	nil,
}

// ServerError
var ServerError Error = Error{
	ErrorInfo{
		Code:    "ServerError",
		Message: "An error occurred because of server exception error.",
	},
	500,
	nil,
}
