package errorobjects

import (
	"fmt"
	"net/http"
)

type ParameterBindingError struct {
	msg            string
	code           int
	HTTPStatusCode int
}

func NewParameterBindingError(msg interface{}) *ParameterBindingError {
	return &ParameterBindingError{msg: fmt.Sprint(msg), code: ErrorCodeParameterBinding, HTTPStatusCode: http.StatusBadRequest}
}

func (pbe *ParameterBindingError) Error() string {
	return pbe.msg + ". error code is " + fmt.Sprint(pbe.code)
}
