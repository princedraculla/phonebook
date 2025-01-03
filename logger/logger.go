package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

// LoadLogger loads logger config
func LoadLogger(logLevel int8) {
	cfg := zap.Config{
		Level:       zap.NewAtomicLevelAt(zapcore.Level(logLevel)),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	Logger = logger.Sugar()
}

var Logg *zap.Logger

func InitializeLogger() *zap.Logger {
	// Create a new production logger
	Logg, err := zap.NewProduction()
	if err != nil {
		panic(err) // Handle the error appropriately in real applications
	}
	return Logg
}
