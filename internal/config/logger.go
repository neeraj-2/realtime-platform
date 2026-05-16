package config

import (
	"go.uber.org/zap"
)

var Log *zap.Logger

func InitLogger() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
	Log = logger
}