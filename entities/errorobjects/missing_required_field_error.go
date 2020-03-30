package errorobjects

import (
	"fmt"
	"net/http"
)

type MissingRequiredFieldsError struct {
	msg            string
	code           int
	HTTPStatusCode int
}

func NewMissingRequiredFieldsError(msg interface{}) *MissingRequiredFieldsError {
	return &MissingRequiredFieldsError{msg: fmt.Sprint(msg), code: ErrorCodeMissingRequiredFiled, HTTPStatusCode: http.StatusBadRequest}
}

func (mrfe *MissingRequiredFieldsError) Error() string {
	return mrfe.msg + ". error code is " + fmt.Sprint(mrfe.code)
}
