package dig

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"gophKeeper/internal/config"
	"gophKeeper/internal/modules/auth/auth_http"
	"gophKeeper/internal/modules/auth/auth_services/auth_service"
	"gophKeeper/internal/modules/pwd/pwd_http"
	"gophKeeper/internal/modules/pwd/pwd_services"
)

func RegistrationServices(
	_ context.Context,
	cfg *config.Config,
	pool *pgxpool.Pool,
) error {

	_ = dg.Provide(func(cfg *config.Config) *config.Config {
		return cfg
	})

	_ = dg.Provide(func(pool *pgxpool.Pool) *auth_service.AuthService {
		return auth_service.NewAuthService(pool)
	})

	_ = dg.Provide(func(auth *auth_service.AuthService) auth_http.AuthHandlers {
		return auth_http.NewAuthHandlersHTTP(auth)
	})

	_ = dg.Provide(func(pool *pgxpool.Pool) *pwd_services.PwdService {
		return pwd_services.NewPwdService(pool)
	})

	_ = dg.Provide(func(pwd *pwd_services.PwdService) pwd_http.PwdHandlers {
		return pwd_http.NewPwdHandlersHTTP(pwd)
	})

	return nil
}
