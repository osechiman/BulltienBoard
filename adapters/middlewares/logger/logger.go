package logger

import (
	"context"
)

// Logger はロギング処理を実装する際に満たすべきインターフェースです。
type Logger interface {
	// Error は予期しないその他の実行時エラーが発生した場合に利用します。
	Error(c *context.Context, msg interface{})
	// Info は実行時の何らかの注目すべき事象を記録したい場合に利用します。（開始や終了など）
	Info(c *context.Context, msg interface{})
	// Debug はシステムの動作状況に関する詳細な情報を記録したい場合に利用します。
	// 暗黙的に本番環境ではこの処理は無視される事を期待しています。
	Debug(c *context.Context, msg interface{})
	// Warn は廃要素となったAPIの使用、APIの不適切な使用、エラーに近い事象などを記録したい場合に利用します。
	// 例えば実行時に生じた異常とは言い切れないが正常とも異なる何らかの予期しない問題が発生した場合などです。
	Warn(c *context.Context, msg interface{})
}
