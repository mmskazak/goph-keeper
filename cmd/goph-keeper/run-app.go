package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"gophKeeper/internal/app"
	"gophKeeper/internal/config"
	"gophKeeper/internal/logger"
	"os"
	"os/signal"
	"syscall"
	"time"
)

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
	logger.Log.Infoln("Приложение запущено.")
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
