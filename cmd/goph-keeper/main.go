package main

import (
	"context"
	"goph-keeper/internal/config"
	"goph-keeper/internal/logger"
	"goph-keeper/internal/storage/psql"
	"log"
	"time"
)

const shutdownDuration = 5 * time.Second

// Главная функция запуска приложения.
func main() {
	ctx := context.Background()

	// Инициализация конфигурации
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalf("Ошибка инициализации конфигурации: %v", err)
	}

	// Инициализация глобального zapLogger
	level, err := cfg.LogLevel.Value()
	if err != nil {
		log.Fatalf("Ошибка получения уровня логирования: %v", err)
	}
	zapLogger, err := logger.InitLogger(level)
	if err != nil {
		log.Printf("Ошибка инициализации глобального логера: %v", err)
	}

	// Инициализируем базу данных
	pool, err := psql.NewPgxPool(ctx, cfg)
	if err != nil {
		zapLogger.Fatalf("Ошибка инициализаци базы данных Postgres: %v", err)
	}

	// Накатываем миграции
	err = psql.RunMigrations(cfg)
	if err != nil {
		zapLogger.Fatalf("Ошибка запуска миграций: %v", err)
	}

	// Запуск приложения
	err = runApp(ctx, cfg, pool, shutdownDuration, zapLogger)
	if err != nil {
		zapLogger.Fatalf("Ошибка запусака приложения: %v", err)
	}
}
