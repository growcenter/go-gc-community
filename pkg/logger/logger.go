package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func Init() {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig= zapcore.EncoderConfig{
		TimeKey:       "time",
        LevelKey:      "level",
        NameKey:       "logger",
        CallerKey:     "caller",
        MessageKey:    "message",
        StacktraceKey: "stacktrace",
        LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel: zapcore.CapitalColorLevelEncoder,
		EncodeTime: zapcore.ISO8601TimeEncoder,
		EncodeCaller: zapcore.FullCallerEncoder,
		EncodeName: zapcore.FullNameEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,

	}
	/*config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	config.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	config.EncoderConfig.EncodeName = zapcore.FullNameEncoder
	config.EncoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder*/

	logger, err := config.Build()
	if err != nil {
		panic("failed to initialize logger")
	}

	Logger = logger
}