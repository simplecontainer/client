package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Log *zap.Logger
var LogFlannel *zap.Logger

func NewLogger(logDir string, logLevel string) *zap.Logger {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	atomicLevel, err := zap.ParseAtomicLevel(logLevel)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	config := zap.Config{
		Level:             atomicLevel,
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			"stderr",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
		InitialFields: map[string]interface{}{},
	}

	return zap.Must(config.Build())
}

func NewLoggerFlannel(logDir string, logLevel string) *zap.Logger {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	atomicLevel, err := zap.ParseAtomicLevel(logLevel)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	flannelStdout := fmt.Sprintf("%s/flannel.log", logDir)
	flannelStderr := fmt.Sprintf("%s/flannel.err", logDir)

	_, err = os.Create(flannelStdout)
	if err != nil {
		panic(err)
	}

	_, err = os.Create(flannelStderr)
	if err != nil {
		panic(err)
	}

	config := zap.Config{
		Level:             atomicLevel,
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			flannelStdout,
		},
		ErrorOutputPaths: []string{
			flannelStderr,
		},
		InitialFields: map[string]interface{}{},
	}

	return zap.Must(config.Build())
}
