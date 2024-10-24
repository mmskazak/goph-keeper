package logger

import (
	"fmt"
	"gophKeeper/internal/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.SugaredLogger

// InitGlobalLogger создает и настраивает новый logger на основе уровня логирования.
func InitGlobalLogger(cfg *config.Config) error {
	// Получение уровня логирования из конфигурации.
	level, err := cfg.LogLevel.Value()
	if err != nil {
		return fmt.Errorf("error getting log level: %w", err)
	}

	// Инициализация logger для установленного уровня
	Log, err = initLogger(level)
	if err != nil {
		return fmt.Errorf("error initializing logger: %w", err)
	}

	return nil
}

func initLogger(level zapcore.Level) (*zap.SugaredLogger, error) {
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
