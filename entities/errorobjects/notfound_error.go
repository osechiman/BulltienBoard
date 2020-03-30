package errorobjects

import (
	"fmt"
	"net/http"
)

type NotFoundError struct {
	msg            string
	code           int
	HTTPStatusCode int
}

func NewNotFoundError(msg interface{}) *NotFoundError {
	return &NotFoundError{msg: fmt.Sprint(msg), code: ErrorCodeNotFound, HTTPStatusCode: http.StatusNotFound}
}

func (nfe *NotFoundError) Error() string {
	return nfe.msg + " not found. error code is " + fmt.Sprint(nfe.code)
}
