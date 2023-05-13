package log

import (
	"backend/src/common"
	"backend/src/common/helpers"
	"context"
)

var globalLogger *logger

func Info(_ context.Context, msg string, args ...interface{}) {
	globalLogger.Info(msg, args...)
}

func Debug(_ context.Context, msg string, args ...interface{}) {
	globalLogger.Debug(msg, args...)
}

func Warn(_ context.Context, msg string, args ...interface{}) {
	globalLogger.Warn(msg, args...)
}

func Error(_ context.Context, msg string, args ...interface{}) {
	globalLogger.Error(msg, args...)
}

func Fatal(msg string, args ...interface{}) {
	globalLogger.Fatal(msg, args...)
}

func IErr(_ context.Context, err *common.Error) {
	if helpers.IsInternalError(err) {
		globalLogger.Error(err.GetDetail())
	} else if helpers.IsClientError(err) {
		globalLogger.Warn(err.ToJSon())
	}

}

func GetLogger() *logger {
	return globalLogger
}
