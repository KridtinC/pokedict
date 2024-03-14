package logger

import "go.uber.org/zap"

func Get() *zap.Logger {
	return zap.L()
}
