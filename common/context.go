package common

import "go.uber.org/zap"

type AppContext struct {
	Logger *zap.Logger
}

func NewAppContext() *AppContext {
	logger, _ := zap.NewProduction()
	return &AppContext{
		Logger: logger,
	}
}
