package goph_keeper

import (
	"context"
	"gophKeeper/internal/config"
	"gophKeeper/internal/logger"
	"gophKeeper/internal/postgres"
	"log"
)

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

	//Инициализируем базу данных
	pool, err := postgres.InitPostgres(ctx, cfg)
	if err != nil {
		logger.Log.Errorf("Ошибка инициализаци базы данных Postgres: %v", err)
	}
	_ = pool

	//Накатываем миграции
	err = postgres.RunMigrations(cfg)
	if err != nil {
		logger.Log.Errorf("Ошибка запуска миграций: %v", err)
	}

	//cfg, zapLog, storage := prepareParamsForApp(ctx)
	//runApp(ctx, cfg, zapLog, storage, shutdownDuration)

	// config.Init()
	// app.StartHTTP()
	// app.StartGRPC
	// app.StopHTTP()
	// app.StopGRPC()
}
