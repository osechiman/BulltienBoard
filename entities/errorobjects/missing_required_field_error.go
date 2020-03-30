package errorobjects

import (
	"fmt"
	"net/http"
)

// MissingRequiredFieldsError は必須項目が指定されていなかった場合に生成するErrorオブジェクトです。
type MissingRequiredFieldsError struct {
	msg            string // msg はError()を呼び出した際に文字列の先頭に結合される文字列です。
	code           int    // code はerror_codeを参照して設定します。
	HTTPStatusCode int    // HTTPStatusCode はHTTPレスポンス時に設定したいレスポンスコードです。
}

// NewMissingRequiredFieldsError はMissingRequiredFieldsErrorオブジェクトを生成します。
func NewMissingRequiredFieldsError(msg interface{}) *MissingRequiredFieldsError {
	return &MissingRequiredFieldsError{msg: fmt.Sprint(msg), code: ErrorCodeMissingRequiredFiled, HTTPStatusCode: http.StatusBadRequest}
}

// Error はエラーメッセージを文字列で返却します。
func (mrfe *MissingRequiredFieldsError) Error() string {
	return mrfe.msg + ". error code is " + fmt.Sprint(mrfe.code)
}
