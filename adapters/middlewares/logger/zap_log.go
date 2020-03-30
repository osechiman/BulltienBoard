package logger

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"vspro/drivers/configs"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// zapLogger はzapインスタンスを再利用する為のパッケージグローバルな変数です。
var zapLogger *zap.Logger

// LogColumns はログに出力したい値をまとめたStructです。
type LogColumns struct {
	PID  string // PID はプロセスIDです。
	UA   string // UA はユーザーエージェントです。
	Path string // Path はリクエストパスです。
	Body []byte // Body はリクエストボディです。

}

// init ではzapのConfigの設定を行います。
func init() {
	var zc zap.Config
	c := configs.GetOsConfigInstance()
	// c := configs.GetYamlConfigInstance()

	switch c.Get().Environment {
	case "production":
		zc = zap.NewProductionConfig()
	default:
		zc = zap.NewDevelopmentConfig()
		zc.Encoding = "json"
	}

	zl, err := zc.Build()
	if err != nil {
		log.Fatal(err.Error())
	}
	zapLogger = zl
	defer zl.Sync()
}

// Error は予期しないその他の実行時エラーが発生した場合に利用します。
func (lc *LogColumns) Error(c *gin.Context, msg interface{}) {
	l := getLoggerColumns(c)
	zapLogger.Error(fmt.Sprint(msg),
		zap.String("processID", l.PID),
		zap.String("user-agent", l.UA),
		zap.String("path", l.Path),
		zap.ByteString("body", l.Body),
	)
}

// Info は実行時の何らかの注目すべき事象を記録したい場合に利用します。（開始や終了など）
func (lc *LogColumns) Info(c *gin.Context, msg interface{}) {
	l := getLoggerColumns(c)
	zapLogger.Info(fmt.Sprint(msg),
		zap.String("processID", l.PID),
		zap.String("user-agent", l.UA),
		zap.String("path", l.Path),
		zap.ByteString("body", l.Body),
	)
}

// Debug はシステムの動作状況に関する詳細な情報を記録したい場合に利用します。
// 暗黙的に本番環境ではこの処理は無視される事を期待しています。
func (lc *LogColumns) Debug(c *gin.Context, msg interface{}) {
	l := getLoggerColumns(c)
	zapLogger.Debug(fmt.Sprint(msg),
		zap.String("processID", l.PID),
		zap.String("user-agent", l.UA),
		zap.String("path", l.Path),
		zap.ByteString("body", l.Body),
	)
}

// Warn は廃要素となったAPIの使用、APIの不適切な使用、エラーに近い事象などを記録したい場合に利用します。
// 例えば実行時に生じた異常とは言い切れないが正常とも異なる何らかの予期しない問題が発生した場合などです。
func (lc *LogColumns) Warn(c *gin.Context, msg interface{}) {
	l := getLoggerColumns(c)
	zapLogger.Warn(fmt.Sprint(msg),
		zap.String("processID", l.PID),
		zap.String("user-agent", l.UA),
		zap.String("path", l.Path),
		zap.ByteString("body", l.Body),
	)
}

// NewZapLogger はzapを利用したロギングStructを作成します。
func NewZapLogger(c *gin.Context) *LogColumns {
	l := &LogColumns{
		PID:  generateProcessID(),
		UA:   getUserAgent(c),
		Path: getPath(c),
		Body: getBody(c),
	}
	return l
}

// getLoggerColumns はコンテキストからLogColumnsを取得します。取得出来ない場合は生成して返却します。
func getLoggerColumns(c *gin.Context) *LogColumns {
	lm, exits := c.Get("Logger")
	if exits {
		return lm.(*LogColumns)
	}

	return &LogColumns{}
}

// generateProcessID はuuidを新規発行して返却します。
func generateProcessID() string {
	uid, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err.Error())
	}
	return uid.String()
}

// getUserAgent はコンテキストからユーザーエージェントを取得します。
func getUserAgent(c *gin.Context) string {
	return c.GetHeader("User-Agent")
}

// getPath はコンテキストからリクエストパスを取得します。
func getPath(c *gin.Context) string {
	return c.Request.URL.Path
}

// getBody はコンテキストからリクエストボディを取得します。
func getBody(c *gin.Context) []byte {
	var buf bytes.Buffer
	tee := io.TeeReader(c.Request.Body, &buf)
	body, err := ioutil.ReadAll(tee)
	if err != nil {
		log.Fatal(err.Error())
	}
	c.Request.Body = ioutil.NopCloser(&buf)
	return body
}

// DefaultLogger はリクエスト開始時と終了時で出力したいログを出力します。
func DefaultLogger(c *gin.Context) {

	zl := NewZapLogger(c)
	// 開始・終了時の共通出力情報生成
	c.Set("Logger", zl)
	// 処理開始時のログ
	zl.Info(c, "start")

	c.Next()
	// 処理終了時のログ
	zl.Info(c, "end")
}
