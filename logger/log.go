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

func Infof(template string, args ...interface{}) {
	l.Infof(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	l.Fatalf(template, args...)
}

func Debugf(template string, args ...interface{}) {
	l.Debugf(template, args...)
}

func Warnf(template string, args ...interface{}) {
	l.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	l.Errorf(template, args...)
}

func Fatal(args ...interface{}) {
	l.Fatal(args)
}

func Debug(args ...interface{}) {
	l.Debug(args)
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
