package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger() {
	conf := zap.Config{
		Level:         zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:   false,
		Encoding:      "console",
		EncoderConfig: zap.NewDevelopmentEncoderConfig(),
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		OutputPaths: []string{
			"stdout",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
	}
	conf.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	l, err := conf.Build()
	if err != nil {
		log.Fatal(err)
	}

	Logger = l
}
