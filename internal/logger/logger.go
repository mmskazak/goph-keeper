package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger(level zapcore.Level) (*zap.SugaredLogger, error) {
	// Создание конфигурации для logger с настройками по умолчанию.
	zapCfg := zap.NewProductionConfig()
	// Установка уровня логирования.
	zapCfg.Level = zap.NewAtomicLevelAt(level)

	// Построение logger на основе конфигурации.
	logger, err := zapCfg.Build()
	if err != nil {
		return nil, fmt.Errorf("ошибка в инициализации логгера %w", err)
	}

	// Создание SugaredLogger для более удобного использования logger с методами Sugared.
	sugar := logger.Sugar()

	return sugar, nil
}
