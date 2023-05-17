package logger

import (
	"github.com/vsPEach/LMS_subsystem/DistributorService/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New(conf config.LoggerConf) *zap.SugaredLogger {
	logConfig := zap.Config{
		Level:            zap.NewAtomicLevel(),
		DisableCaller:    false,
		Development:      true,
		Encoding:         "console",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "T",
			LevelKey:       "L",
			NameKey:        "N",
			CallerKey:      "C",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "M",
			StacktraceKey:  "S",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}
	return zap.Must(logConfig.Build()).Sugar()
}
