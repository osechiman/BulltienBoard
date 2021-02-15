package presenters

// HTTPResponse はHTTPプロトコルのリクエストに対する共通のレスポンスです。
// HTTPプロトコルに対するレスポンスの場合は全てこのStructを利用してください。
type HTTPResponse struct {
	Items interface{}
}

// newHTTPResponse はHTTPプロトコルに対するStructを返却します。
func newHTTPResponse(data interface{}) *HTTPResponse {
	hr := HTTPResponse{
		Items: data,
	}
	return &hr
}

// newHTTPResponse は正常時のHTTPプロトコルに対するStructを返却します。
func newHTTPSuccessResponse(i interface{}) *HTTPResponse {
	return newHTTPResponse(i)
}

// newHTTPResponse はエラー時のHTTPプロトコルに対するStructを返却します。
func newHTTPErrorResponse(err error) *HTTPResponse {
	return newHTTPResponse(err)
}
