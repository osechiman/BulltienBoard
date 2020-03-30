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

var logger *zap.Logger

type LogColumns struct {
	PID  string
	UA   string
	Path string
	Body []byte
}

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
	logger = zl
	defer zl.Sync()
}

func Error(c *gin.Context, msg interface{}) {
	lm := getLoggerColumns(c)
	logger.Error(fmt.Sprint(msg),
		zap.String("processID", lm.PID),
		zap.String("user-agent", lm.UA),
		zap.String("path", lm.Path),
		zap.ByteString("body", lm.Body),
	)
}

func Info(c *gin.Context, msg interface{}) {
	lm := getLoggerColumns(c)
	logger.Info(fmt.Sprint(msg),
		zap.String("processID", lm.PID),
		zap.String("user-agent", lm.UA),
		zap.String("path", lm.Path),
		zap.ByteString("body", lm.Body),
	)
}

func Debug(c *gin.Context, msg interface{}) {
	lm := getLoggerColumns(c)
	logger.Debug(fmt.Sprint(msg),
		zap.String("processID", lm.PID),
		zap.String("user-agent", lm.UA),
		zap.String("path", lm.Path),
		zap.ByteString("body", lm.Body),
	)
}

func Warn(c *gin.Context, msg interface{}) {
	lm := getLoggerColumns(c)
	logger.Warn(fmt.Sprint(msg),
		zap.String("processID", lm.PID),
		zap.String("user-agent", lm.UA),
		zap.String("path", lm.Path),
		zap.ByteString("body", lm.Body),
	)
}

func NewLoggerMan(c *gin.Context) *LogColumns {
	lm := &LogColumns{
		PID:  generateProcessID(),
		UA:   getUserAgent(c),
		Path: getPath(c),
		Body: getBody(c),
	}
	return lm
}

func getLoggerColumns(c *gin.Context) *LogColumns {
	lm, exits := c.Get("Logger")
	if exits {
		return lm.(*LogColumns)
	}

	return &LogColumns{}
}

func generateProcessID() string {
	uid, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err.Error())
	}
	return uid.String()
}

func getUserAgent(c *gin.Context) string {
	return c.GetHeader("User-Agent")
}

func getPath(c *gin.Context) string {
	return c.Request.URL.Path
}

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

func DefaultLogger(c *gin.Context) {

	// 開始・終了時の共通出力情報生成
	c.Set("Logger", NewLoggerMan(c))
	// 処理開始時のログ
	Info(c, "start")

	c.Next()
	// 処理終了時のログ
	Info(c, "end")

}
