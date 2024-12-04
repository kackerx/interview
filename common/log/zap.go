package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _logger *zap.Logger

func init() {
	encoderConfg := zap.NewProductionEncoderConfig()
	encoderConfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewJSONEncoder(encoderConfg)

	var cores []zapcore.Core
	cores = append(
		cores,
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.DebugLevel),
	)

	core := zapcore.NewTee(cores...)
	_logger = zap.New(core)
}
