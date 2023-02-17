package logs

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger
var err error

func init() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.StacktraceKey = ""
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}
func Error(message interface{}, fields ...zap.Field) {
	switch v := message.(type) {
	case error:
		log.Error(v.Error(), fields...)
	case string:
		log.Error(v, fields...)
	}
}
func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}
