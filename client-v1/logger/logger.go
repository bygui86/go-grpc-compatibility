package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/bygui86/go-grpc-compatibility/client-v1/utils"
)

const (
	logEncodingEnvVar = "LOG_ENCODING"
	logLevelEnvVar    = "LOG_LEVEL"

	logEncodingDefault = "console"
	logLevelDefault    = "info"
)

var SugaredLogger *zap.SugaredLogger

func init() {
	encoding := utils.GetString(logEncodingEnvVar, logEncodingDefault)
	level := utils.GetString(logLevelEnvVar, logLevelDefault)
	zapLevel := zapcore.InfoLevel
	err := zapLevel.Set(level)
	if err != nil {
		fmt.Printf("Error initializing zap logger: %s\n", err.Error())
		os.Exit(1)
	}
	buildLogger(encoding, zapLevel)
}

func buildLogger(encoding string, level zapcore.Level) {
	logger, err := zap.Config{
		Encoding:         encoding,
		Level:            zap.NewAtomicLevelAt(level),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig:    buildEncoderConfig(),
	}.Build()
	if err != nil {
		fmt.Printf("Error building zap logger: %s\n", err.Error())
		os.Exit(1)
	}
	SugaredLogger = logger.Sugar()
}

func buildEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:   "message",
		TimeKey:      "time",
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		LevelKey:     "level",
		EncodeLevel:  zapcore.CapitalLevelEncoder,
		CallerKey:    "caller",
		EncodeCaller: zapcore.ShortCallerEncoder,
	}
}
