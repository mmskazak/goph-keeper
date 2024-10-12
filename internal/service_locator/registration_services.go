package service_locator

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"gophKeeper/internal/config"
	"gophKeeper/internal/modules/auth/http"
	"gophKeeper/internal/modules/auth/services/auth_service"
)

func RegistrationServices(
	_ context.Context,
	cfg *config.Config,
	pool *pgxpool.Pool,
) {
	sc := InitServiceLocator()
	sc.Register("config", cfg)
	sc.Register("pool", pool)

	authService := auth_service.NewAuthService(pool)
	sc.Register("auth_service", authService)

	authHandlersHTTP := auth_http.NewAuthHandlersHTTP(authService)
	sc.Register("auth_handlers_http", authHandlersHTTP)
}
