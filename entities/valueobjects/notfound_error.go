package valueobjects

import "fmt"

type NotFoundError struct {
	msg  string
	code string
}

func NewNotFoundError(msg interface{}) *NotFoundError {
	return &NotFoundError{msg: fmt.Sprint(msg), code: "001"}
}

func (nfe *NotFoundError) Error() string {
	return nfe.msg + " is not found. error code is " + nfe.code
}
