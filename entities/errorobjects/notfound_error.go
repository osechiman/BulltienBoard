package errorobjects

import (
	"fmt"
)

// NotFoundError は取得しようとした値が存在しなかった場合に生成するErrorオブジェクトです。
type NotFoundError struct {
	msg  string // msg はError()を呼び出した際に文字列の先頭に結合される文字列です。
	code int    // code はerror_codeを参照して設定します。
}

// NewNotFoundError はNotFoundErrorオブジェクトを生成します。
func NewNotFoundError(msg interface{}) *NotFoundError {
	return &NotFoundError{msg: fmt.Sprint(msg), code: ErrorCodeNotFound}
}

// Error はエラーメッセージを文字列で返却します。
func (nfe *NotFoundError) Error() string {
	return nfe.msg + " not found. error code is " + fmt.Sprint(nfe.code)
}
