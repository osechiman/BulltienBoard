package errorobjects

import (
	"fmt"
)

// ResourceLimitedError は保存可能なデータ件数の閾値を超えている場合に生成するErrorオブジェクトです。
type ResourceLimitedError struct {
	msg  string // msg はError()を呼び出した際に文字列の先頭に結合される文字列です。
	code int    // code はerror_codeを参照して設定します。
}

// NewResourceLimitedError はResourceLimitedErrorオブジェクトを生成します。
func NewResourceLimitedError(msg interface{}) *ResourceLimitedError {
	return &ResourceLimitedError{msg: fmt.Sprint(msg), code: ErrorCodeResourceLimitedError}
}

// Error はエラーメッセージを文字列で返却します。
func (rle *ResourceLimitedError) Error() string {
	return rle.msg + ". error code is " + fmt.Sprint(rle.code)
}
