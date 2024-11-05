package app

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"goph-keeper/internal/config"
	"goph-keeper/internal/modules/auth/authgrpc"
	"goph-keeper/internal/modules/auth/authservices/authservice"
	"goph-keeper/internal/modules/auth/proto"
	"goph-keeper/internal/modules/file/filegrpc"
	"goph-keeper/internal/modules/file/fileservices"
	proto3 "goph-keeper/internal/modules/file/proto"
	proto2 "goph-keeper/internal/modules/pwd/proto"
	"goph-keeper/internal/modules/pwd/pwdgrpc"
	"goph-keeper/internal/modules/pwd/pwdservices"
	"net/http"
	"time"

	"net"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/acme/autocert"
	"google.golang.org/grpc"
)

// Константы для таймаутов.
const readTimeout = 5 * time.Second
const writeTimeout = 5 * time.Second

// App представляет приложение с HTTP и gRPC серверами.
type App struct {
	httpServer *http.Server
	grpcServer *grpc.Server
	zapLogger  *zap.SugaredLogger
}

// NewApp создает новый экземпляр приложения с HTTP и gRPC серверами.
func NewApp(
	ctx context.Context,
	cfg *config.Config,
	pool *pgxpool.Pool,
	zapLogger *zap.SugaredLogger,
) *App {

	pwdServices := pwdservices.NewPwdService(pool, cfg.EncryptionKey, zapLogger)

	fileService := fileservices.NewFileService(
		pool,
		cfg.EncryptionKey,
		cfg.MaxFileSize,
		zapLogger,
	)

	router := chi.NewRouter()
	router = registrationHandlersHTTP(ctx, router, cfg, pool)

	manager := &autocert.Manager{
		// перечень доменов, для которых будут поддерживаться сертификаты
		HostPolicy: autocert.HostWhitelist("localhost"),
	}

	httpServer := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		// для TLS-конфигурации используем менеджер сертификатов
		TLSConfig: manager.TLSConfig(),
	}

	grpcServer := grpc.NewServer()
	// Регистрация gRPC-сервисов
	// Регистрация обработчиков для аутентификации
	authGRPSHandlers := authgrpc.NewAuthGRPCServer(
		authservice.NewAuthService(pool),
		cfg.SecretKey,
		zapLogger,
	)
	proto.RegisterAuthServiceServer(grpcServer, authGRPSHandlers) // Регистрация сервиса аутентификации
	// Регистрация обработчиков для сохранения паролей
	passwordGRPCHandlers := pwdgrpc.NewPasswordGRPCServer(
		pwdServices,
		cfg.SecretKey,
		zapLogger,
	)
	proto2.RegisterPasswordServiceServer(grpcServer, passwordGRPCHandlers) // Регистрация сервиса паролей
	// Регистрация обработчиков для сохранения файлов
	fileGRPCHandlers := filegrpc.NewFileGRPCServer(
		fileService,
		cfg.SecretKey,
		zapLogger,
	)
	proto3.RegisterFileServiceServer(grpcServer, fileGRPCHandlers)

	return &App{
		httpServer: httpServer,
		grpcServer: grpcServer,
		zapLogger:  zapLogger,
	}
}

// Start запускает HTTP и gRPC сервера.
func (a *App) Start() error {
	a.zapLogger.Infoln("Приложение запущено.")

	// Запуск HTTP сервера в отдельной горутине
	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.zapLogger.Errorf("error listen and serve: %v", err)
		}
	}()

	// Запуск gRPC сервера
	listener, err := net.Listen("tcp", ":50051") // Порт для gRPC
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	go func() {
		if err := a.grpcServer.Serve(listener); err != nil {
			a.zapLogger.Errorf("gRPC server failed: %v", err)
		}
	}()

	return nil
}

// Stop корректно завершает работу приложения.
func (a *App) Stop(ctx context.Context) error {
	// Закрытие gRPC сервера
	a.grpcServer.GracefulStop()

	// Закрытие HTTP сервера с учетом переданного контекста.
	if err := a.httpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("err Shutdown server: %w", err)
	}

	return nil
}
