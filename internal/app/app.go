package app

import (
	"context"
	"errors"
	"fmt"
	"gophKeeper/internal/config"
	"gophKeeper/internal/logger"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/acme/autocert"
)

const readTimeout = 5 * time.Second
const writeTimeout = 5 * time.Second

// App представляет приложение с HTTP сервером и логгером.
type App struct {
	server *http.Server
}

// NewApp создает новый экземпляр приложения.
func NewApp(
	ctx context.Context,
	cfg *config.Config,
	pool *pgxpool.Pool,
) *App {
	router := chi.NewRouter()
	router = registrationHandlersHTTP(ctx, router, cfg, pool)

	manager := &autocert.Manager{
		// перечень доменов, для которых будут поддерживаться сертификаты
		HostPolicy: autocert.HostWhitelist("localhost"),
	}

	return &App{
		server: &http.Server{
			Addr:         cfg.Address,
			Handler:      router,
			ReadTimeout:  readTimeout,
			WriteTimeout: writeTimeout,
			// для TLS-конфигурации используем менеджер сертификатов
			TLSConfig: manager.TLSConfig(),
		},
	}
}

// Start запускает HTTP сервер.
func (a *App) Start() error {
	logger.Log.Infoln("Приложение запущено.")
	err := a.server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("error listen and serve: %w", err)
	}
	return nil
}

// Stop корректно завершает работу приложения.
func (a *App) Stop(ctx context.Context) error {
	// Закрытие сервера с учетом переданного контекста.
	if err := a.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("err Shutdown server: %w", err)
	}
	return nil
}
