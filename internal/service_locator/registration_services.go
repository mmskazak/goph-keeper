package service_locator

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"gophKeeper/internal/config"
	auth_http "gophKeeper/internal/modules/auth/auth_http"
	"gophKeeper/internal/modules/auth/auth_services/auth_service"
	pwd_http "gophKeeper/internal/modules/pwd/pwd_http"
	pwd_services "gophKeeper/internal/modules/pwd/pwd_services"
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

	pwdService := pwd_services.NewPwdService(pool)
	sc.Register("pwd_service", pwdService)

	pwd_http.NewAuthHandlersHTTP(pwdService)
	sc.Register("pwd_handlers_http", pwd_http.NewAuthHandlersHTTP(pwdService))
}
