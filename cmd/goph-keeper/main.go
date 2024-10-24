package main

import (
	"context"
	"gophKeeper/internal/config"
	"gophKeeper/internal/logger"
	"gophKeeper/internal/storage/psql"
	"log"
	"time"
)

const shutdownDuration = 5 * time.Second

func main() {
	ctx := context.Background()

	// Инициализация конфигурации.
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalf("Ошибка инициализации конфигурации: %v", err)
	}

	// Инициализация глобального logger.
	err = logger.InitGlobalLogger(cfg)
	if err != nil {
		log.Printf("Ошибка инициализации глобального логера: %v", err)
	}

	// Инициализируем базу данных
	pool, err := psql.NewPgxPool(ctx, cfg)
	if err != nil {
		logger.Log.Fatalf("Ошибка инициализаци базы данных Postgres: %v", err)
	}

	// Накатываем миграции
	err = psql.RunMigrations(cfg)
	if err != nil {
		logger.Log.Fatalf("Ошибка запуска миграций: %v", err)
	}

	err = runApp(ctx, cfg, pool, shutdownDuration)
	if err != nil {
		logger.Log.Fatalf("Ошибка запусака приложения: %v", err)
	}
}
