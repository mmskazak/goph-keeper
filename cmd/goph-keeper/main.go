package goph_keeper

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"gophKeeper/internal/app"
	"gophKeeper/internal/config"
	"gophKeeper/internal/logger"
	"gophKeeper/internal/postgres"
	"log"
	"os"
	"os/signal"
	"syscall"
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

	err = runApp(ctx, cfg, pool, shutdownDuration)
	if err != nil {
		logger.Log.Errorf("Ошибка запусака приложения: %v", err)
	}
}

func runApp(
	ctx context.Context,
	cfg *config.Config,
	pool *pgxpool.Pool,
	shutdownDuration time.Duration,
) error {
	defer func() {
		pool.Close()
	}()

	newApp := app.NewApp(ctx, cfg)
	err := newApp.Start()
	if err != nil {
		return fmt.Errorf("error start app: %w", err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case <-quit: // Ожидание сигнала завершения
		logger.Log.Infoln("Получен сигнал завершения, остановка сервера...")
	case <-ctx.Done(): // Завершение по контексту
		logger.Log.Infoln("Контекст завершён, остановка сервера...")
	}

	ctxShutdown, cancel := context.WithTimeout(context.Background(), shutdownDuration)
	defer cancel()

	if err := newApp.Stop(ctxShutdown); err != nil {
		return fmt.Errorf("error stop app: %w", err)
	}
	logger.Log.Infoln("Приложение завершило работу.")

	return nil
}
