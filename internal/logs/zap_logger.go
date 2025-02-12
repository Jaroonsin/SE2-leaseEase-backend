package logs

import (
	"go.uber.org/zap"
)

const (
	DEV  = "development"
	PROD = "production"
)

func NewLogger() *zap.Logger {
	return newLoggerFactory(DEV)
}

func newLoggerFactory(env string) *zap.Logger {
	switch env {
	case DEV:
		return zap.Must(zap.NewDevelopment())
	case PROD:
		return zap.Must(zap.NewProduction())
	default:
		return nil
	}
}
