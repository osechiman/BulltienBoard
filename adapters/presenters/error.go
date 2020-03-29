package presenters

import (
	"net/http"
)

type ErrorPresenter struct{}

func NewErrorPresenter() *ErrorPresenter {
	return &ErrorPresenter{}
}

func (ep *ErrorPresenter) ConvertToHttpErrorResponse(httpStatusCode int, err error) *HTTPResponse {
	return newHTTPErrorResponse(httpStatusCode, http.StatusText(httpStatusCode), err)
}
