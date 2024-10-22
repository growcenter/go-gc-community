package logrus

import "github.com/sirupsen/logrus"

type Logger interface {
	Debug(msg string, params map[string]interface{})
	Info(msg string, params map[string]interface{})
	Warn(msg string, params map[string]interface{})
	Error(msg string, params map[string]interface{})
}

func Debug(msg ...interface{}) {
	logrus.Debug(msg...)
}

func Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

func Info(msg ...interface{}) {
	logrus.Info(msg...)
}

func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func Warn(msg ...interface{}) {
	logrus.Warn(msg...)
}

func Warnf(format string, args ...interface{}) {
	logrus.Warnf(format, args...)
}

func Error(msg ...interface{}) {
	logrus.Error(msg...)
}

func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}