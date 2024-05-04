package logging

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Sugar *zap.SugaredLogger

func InitLogger(debug bool) {
	logFile, err := os.OpenFile("sysmonmq.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		log.Fatal("Error opening log file:", err)
		os.Exit(1)
	}

	var loggerConfig zap.Config

	if debug {
		loggerConfig = zap.NewDevelopmentConfig()
	} else {
		loggerConfig = zap.NewProductionConfig()
		loggerConfig.Encoding = "console"
		loggerConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		loggerConfig.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	}

	loggerConfig.OutputPaths = []string{logFile.Name()}
	loggerConfig.ErrorOutputPaths = []string{logFile.Name()}

	logger, err := loggerConfig.Build()
	if err != nil {
		log.Fatalf("Error building zap logger. %v", err)
	}

	defer logger.Sync()

	Sugar = logger.Sugar()
}

func Debug(args ...interface{}) {
	Sugar.WithOptions(zap.AddCallerSkip(1)).Debug(args...)
}

func Info(args ...interface{}) {
	Sugar.WithOptions(zap.AddCallerSkip(1)).Info(args...)
}

func Warn(args ...interface{}) {
	Sugar.WithOptions(zap.AddCallerSkip(1)).Warn(args...)
}

func Error(args ...interface{}) {
	Sugar.WithOptions(zap.AddCallerSkip(1)).Error(args...)
}

func Fatal(args ...interface{}) {
	Sugar.WithOptions(zap.AddCallerSkip(1)).Fatal(args...)
	os.Exit(1)
}
