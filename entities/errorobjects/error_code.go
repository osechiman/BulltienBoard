package errorobjects

const (
	ErrorCodeInternalServerError     int = iota // 原因不明のエラー level:Error
	ErrorCodeNotFound                           // 取得しようとした値が存在しなかった level:Debug
	ErrorCodeMissingRequiredFiled               // 必須項目が指定されていなかった level:Warn
	ErrorCodeParameterBinding                   // リクエストパラメータをStructにバインドしようとした結果失敗した level:Warn
	ErrorCodeCharacterSizeValidation            // 文字サイズが仕様通りでは無い level:Warn
	ErrorCodeResourceLimitedError               // 保存可能なデータ件数の閾値を超えている level:Warn
)
