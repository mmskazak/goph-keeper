package main

import (
	"context"
	"fmt"
	"gophKeeper/internal/app"
	"gophKeeper/internal/dig"
	"gophKeeper/internal/logger"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func runApp(
	ctx context.Context,
	shutdownDuration time.Duration,
) error {
	cfg, err := dig.GetConfig()
	if err != nil {
		logger.Log.Errorf("Error getting config: %v", err)
	}

	newApp := app.NewApp(ctx, cfg)
	go func() {
		err := newApp.Start()
		if err != nil {
			logger.Log.Errorf("error start app: %v", err)
		}
	}()

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
