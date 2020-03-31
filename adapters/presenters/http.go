package presenters

// HTTPResponse はHTTPプロトコルのリクエストに対する共通のレスポンスです。
// HTTPプロトコルに対するレスポンスの場合は全てこのStructを利用してください。
type HTTPResponse struct {
	Status  int         // Status はHTTPステータスコードです。多くの場合はhttpパッケージの定数を使います。
	Message string      // Message はHTTPステータスコードに紐付いたメッセージです。多くの場合はhttpパッケージの定数を使います。
	Data    interface{} // Data はレスポンスデータを格納します。
}

// newHTTPResponse はHTTPプロトコルに対するStructを返却します。
func newHTTPResponse(status int, message string, data interface{}) *HTTPResponse {
	return &HTTPResponse{Status: status, Message: message, Data: data}
}

// newHTTPResponse は正常時のHTTPプロトコルに対するStructを返却します。
func newHTTPSuccessResponse(status int, message string, i interface{}) *HTTPResponse {
	res := make([]interface{}, 0)
	res = append(res, i)
	return newHTTPResponse(status, message, res)
}

// newHTTPResponse はエラー時のHTTPプロトコルに対するStructを返却します。
func newHTTPErrorResponse(status int, message string, err error) *HTTPResponse {
	res := make([]string, 0)
	res = append(res, err.Error())
	return newHTTPResponse(status, message, res)
}
