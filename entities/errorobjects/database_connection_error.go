package errorobjects

import (
	"fmt"
)

// DatabaseConnectionError は原因不明のエラーの場合に生成するErrorオブジェクトです。
type DatabaseConnectionError struct {
	msg  string // msg はError()を呼び出した際に文字列の先頭に結合される文字列です。
	code int    // code はerror_codeを参照して設定します。
}

// NewDatabaseConnectionError はDatabaseConnectionErrorオブジェクトを生成します。
func NewDatabaseConnectionError(msg interface{}) *DatabaseConnectionError {
	return &DatabaseConnectionError{msg: fmt.Sprint(msg), code: ErrorCodeDatabaseConnectionError}
}

// Error はエラーメッセージを文字列で返却します。
func (dce *DatabaseConnectionError) Error() string {
	return dce.msg + ". error code is " + fmt.Sprint(dce.code)
}
