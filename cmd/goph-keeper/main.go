package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"gophKeeper/internal/config"
	"gophKeeper/internal/logger"
	servive_http "gophKeeper/internal/modules/auth/http"
	auth_service "gophKeeper/internal/modules/auth/services/auth_service"
	"gophKeeper/internal/service_locator"
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

	//Инициализируем базу данных
	pool, err := psql.NewPgxPool(ctx, cfg)
	if err != nil {
		logger.Log.Fatalf("Ошибка инициализаци базы данных Postgres: %v", err)
	}

	//Накатываем миграции
	err = psql.RunMigrations(cfg)
	if err != nil {
		logger.Log.Fatalf("Ошибка запуска миграций: %v", err)
	}

	//Регистрация всех сервисов приложения
	registrationServices(ctx, cfg, pool)

	err = runApp(ctx, shutdownDuration)
	if err != nil {
		logger.Log.Fatalf("Ошибка запусака приложения: %v", err)
	}
}

func registrationServices(
	_ context.Context,
	cfg *config.Config,
	pool *pgxpool.Pool,
) {
	sc := service_locator.InitServiceLocator()
	sc.Register("config", cfg)
	sc.Register("pool", pool)

	authService := auth_service.NewAuthService(pool)
	sc.Register("auth_service", authService)

	authHandlersHTTP := servive_http.NewAuthHandlersHTTP(authService)
	sc.Register("auth_handlers_http", authHandlersHTTP)
}
