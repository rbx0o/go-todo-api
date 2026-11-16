package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New() *zap.Logger {
	cfg := zap.NewDevelopmentEncoderConfig()
	enc := zapcore.NewConsoleEncoder(cfg)

	ws := zapcore.AddSync(os.Stdout)

	core := zapcore.NewCore(enc, ws, zapcore.InfoLevel)

	return zap.New(
		core,
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
}
