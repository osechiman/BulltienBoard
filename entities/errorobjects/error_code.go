package errorobjects

const (
	// 原因不明のエラー level:Error
	ErrorCodeInternalServerError int = iota
	// 取得しようとした値が存在しなかった level:Debug
	ErrorCodeNotFound
	// 必須項目が指定されていなかった level:Warn
	ErrorCodeMissingRequiredFiled
	// リクエストパラメータをStructにバインドしようとした結果失敗した level:Warn
	ErrorCodeParameterBinding
)
