package errorobjects

const (
	// ErrorCodeInternalServerError は原因不明のエラー level:Error
	ErrorCodeInternalServerError int = iota
	// ErrorCodeNotFound は取得しようとした値が存在しなかった level:Debug
	ErrorCodeNotFound
	// ErrorCodeMissingRequiredFiled は必須項目が指定されていなかった level:Warn
	ErrorCodeMissingRequiredFiled
	// ErrorCodeParameterBinding はリクエストパラメータをStructにバインドしようとした結果失敗した level:Warn
	ErrorCodeParameterBinding
	// ErrorCodeCharacterSizeValidation は文字サイズが仕様通りでは無い level:Warn
	ErrorCodeCharacterSizeValidation
	// ErrorCodeResourceLimitedError は保存可能なデータ件数の閾値を超えている level:Warn
	ErrorCodeResourceLimitedError
	// ErrorCodeDatabaseConnectionError はデータベース接続ができていない level:Error
	ErrorCodeDatabaseConnectionError
)
