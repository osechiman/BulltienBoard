package errorobjects

import (
	"fmt"
)

// InternalServerError は原因不明のエラーの場合に生成するErrorオブジェクトです。
type InternalServerError struct {
	msg  string // msg はError()を呼び出した際に文字列の先頭に結合される文字列です。
	code int    // code はerror_codeを参照して設定します。
}

// NewInternalServerError はInternalServerErrorオブジェクトを生成します。
func NewInternalServerError(msg interface{}) *InternalServerError {
	return &InternalServerError{msg: fmt.Sprint(msg), code: ErrorCodeInternalServerError}
}

// Error はエラーメッセージを文字列で返却します。
func (ise *InternalServerError) Error() string {
	return ise.msg + ". error code is " + fmt.Sprint(ise.code)
}
