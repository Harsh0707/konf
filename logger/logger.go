package logger

import (
	"fmt"
	"os"

	"bygui86/konf/config/envvar"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// TODO move to konf.go flags?
const (
	logEncodingEnvVar = "KONF_LOG_ENCODING"
	logLevelEnvVar    = "KONF_LOG_LEVEL"

	logEncodingDefault = "console"
	logLevelDefault    = "info"
)

var Logger *zap.Logger
var SugaredLogger *zap.SugaredLogger

func init() {
	encoding := envvar.GetString(logEncodingEnvVar, logEncodingDefault)
	levelString := envvar.GetString(logLevelEnvVar, logLevelDefault)
	level := zapcore.InfoLevel
	err := level.Set(levelString)
	if err != nil {
		fmt.Printf("❌ Error initializing zap logger: %s\n", err.Error())
		os.Exit(1)
	}
	buildLogger(encoding, level)
}

func buildLogger(encoding string, level zapcore.Level) {
	Logger, _ = zap.Config{
		Encoding:         encoding,
		Level:            zap.NewAtomicLevelAt(level),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig:    buildEncoderConfig(level),
	}.Build()
	SugaredLogger = Logger.Sugar()
}

func buildEncoderConfig(level zapcore.Level) zapcore.EncoderConfig {
	if level == zapcore.DebugLevel {
		return zapcore.EncoderConfig{
			MessageKey:   "message",
			TimeKey:      "time",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			LevelKey:     "level",
			EncodeLevel:  zapcore.CapitalLevelEncoder,
			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		}
	} else {
		return zapcore.EncoderConfig{
			MessageKey: "message",
		}
	}
}
