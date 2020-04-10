package errorobjects

import (
	"fmt"
)

// CharacterSizeValidationError は文字サイズが仕様通りでは無い場合に生成するErrorオブジェクトです。
type CharacterSizeValidationError struct {
	msg  string // msg はError()を呼び出した際に文字列の先頭に結合される文字列です。
	code int    // code はerror_codeを参照して設定します。
}

// NewCharacterSizeValidationError はCharacterSizeValidationErrorオブジェクトを生成します。
func NewCharacterSizeValidationError(msg interface{}) *CharacterSizeValidationError {
	return &CharacterSizeValidationError{msg: fmt.Sprint(msg), code: ErrorCodeCharacterSizeValidation}
}

// Error はエラーメッセージを文字列で返却します。
func (csve *CharacterSizeValidationError) Error() string {
	return csve.msg + ". error code is " + fmt.Sprint(csve.code)
}
