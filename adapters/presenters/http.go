package presenters

type HTTPResponse struct {
	Status  int
	Message string
	Data    interface{}
}

func newHTTPResponse(status int, message string, data interface{}) *HTTPResponse {
	return &HTTPResponse{Status: status, Message: message, Data: data}
}

func newHTTPSuccessResponse(status int, message string, i interface{}) *HTTPResponse {
	res := make([]interface{}, 0)
	res = append(res, i)
	return newHTTPResponse(status, message, res)
}

func newHTTPErrorResponse(status int, message string, err error) *HTTPResponse {
	res := make([]string, 0)
	res = append(res, err.Error())
	return newHTTPResponse(status, message, res)
}
