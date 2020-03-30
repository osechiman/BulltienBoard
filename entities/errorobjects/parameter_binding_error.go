package errorobjects

import (
	"fmt"
	"net/http"
)

// ParameterBindingError はリクエストパラメータをStructにバインドしようとした結果失敗した場合に生成するErrorオブジェクトです。
type ParameterBindingError struct {
	msg            string // msg はError()を呼び出した際に文字列の先頭に結合される文字列です。
	code           int    // code はerror_codeを参照して設定します。
	HTTPStatusCode int    // HTTPStatusCode はHTTPレスポンス時に設定したいレスポンスコードです。
}

// NewParameterBindingError はParameterBindingErrorオブジェクトを生成します。
func NewParameterBindingError(msg interface{}) *ParameterBindingError {
	return &ParameterBindingError{msg: fmt.Sprint(msg), code: ErrorCodeParameterBinding, HTTPStatusCode: http.StatusBadRequest}
}

// Error はエラーメッセージを文字列で返却します。
func (pbe *ParameterBindingError) Error() string {
	return pbe.msg + ". error code is " + fmt.Sprint(pbe.code)
}
