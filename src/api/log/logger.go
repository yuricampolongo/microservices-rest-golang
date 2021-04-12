package log

import "go.uber.org/zap"

var (
	logger *zap.Logger
)

func init() {
	newLogger, _ := zap.NewProduction()
	defer newLogger.Sync()
	logger = newLogger
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}
