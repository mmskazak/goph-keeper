package config

import (
	"errors"
	"go.uber.org/zap/zapcore"
	"strings"
)

// LogLevel представляет уровень логирования.
type LogLevel string

// Value возвращает уровень логирования в формате zapcore. Level.
func (ll LogLevel) Value() (zapcore.Level, error) {
	switch strings.ToLower(string(ll)) {
	case "debug":
		return zapcore.DebugLevel, nil
	case "info":
		return zapcore.InfoLevel, nil
	case "warn", "warning":
		return zapcore.WarnLevel, nil
	case "error":
		return zapcore.ErrorLevel, nil
	case "dpanic":
		return zapcore.DPanicLevel, nil
	case "panic":
		return zapcore.PanicLevel, nil
	case "fatal":
		return zapcore.FatalLevel, nil
	default:
		return zapcore.DebugLevel, errors.New("не найдено соответствие текстовому значению LogLevel, " +
			"уровень логирования задан debug")
	}
}
