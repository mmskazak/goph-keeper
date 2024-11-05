package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"goph-keeper/internal/app"
	"goph-keeper/internal/config"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// runApp запуск приложения.
func runApp(
	ctx context.Context,
	cfg *config.Config,
	pool *pgxpool.Pool,
	shutdownDuration time.Duration,
	zapLogger *zap.SugaredLogger,
) error {
	newApp := app.NewApp(ctx, cfg, pool, zapLogger)
	go func() {
		err := newApp.Start()
		if err != nil {
			zapLogger.Errorf("error start app: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case <-quit: // Ожидание сигнала завершения
		zapLogger.Infoln("Получен сигнал завершения, остановка сервера...")
	case <-ctx.Done(): // Завершение по контексту
		zapLogger.Infoln("Контекст завершён, остановка сервера...")
	}

	ctxShutdown, cancel := context.WithTimeout(context.Background(), shutdownDuration)
	defer cancel()

	if err := newApp.Stop(ctxShutdown); err != nil {
		return fmt.Errorf("error stop app: %w", err)
	}
	zapLogger.Infoln("Приложение завершило работу.")

	return nil
}
