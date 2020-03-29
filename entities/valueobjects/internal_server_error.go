package valueobjects

import (
	"fmt"
	"net/http"
)

type InternalServerError struct {
	msg            string
	code           int
	HTTPStatusCode int
}

func NewInternalServerError(msg interface{}) *InternalServerError {
	return &InternalServerError{msg: fmt.Sprint(msg), code: ErrorCodeInternalServerError, HTTPStatusCode: http.StatusInternalServerError}
}

func (ise *InternalServerError) Error() string {
	return ise.msg + ". error code is " + fmt.Sprint(ise.code)
}
