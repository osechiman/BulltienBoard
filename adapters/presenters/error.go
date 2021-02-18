package presenters

// ErrorPresenter はエラー時に外部へ渡す為にデータの変換を行います。
type ErrorPresenter struct{}

// NewErrorPresenter はErrorPresenterを初期化します。
func NewErrorPresenter() *ErrorPresenter {
	return &ErrorPresenter{}
}

// ConvertToHTTPErrorResponse はエラー時のレスポンスを返却します。
func (ep *ErrorPresenter) ConvertToHTTPErrorResponse(err error) *HTTPResponse {
	return newHTTPErrorResponse(err)
}
