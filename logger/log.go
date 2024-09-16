package logger

import (
	"go.uber.org/zap"
)

var l *zap.SugaredLogger

func NewLogger() {
	prod, _ := zap.NewProduction()
	defer prod.Sync() // flushes buffer, if any
	l = prod.Sugar()
}

func init() {
	NewLogger()
}

func Fatal(args ...interface{}) {
	l.Fatal(args)
}

func Info(args ...interface{}) {
	l.Info(args)
}

func Warn(args ...interface{}) {
	l.Warn(args)
}

func Error(args ...interface{}) {
	l.Error(args)
}
