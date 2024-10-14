package dig

import (
	"context"
	"fmt"
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
	err := RegisterConfig(*cfg)
	if err != nil {
		return fmt.Errorf("error setting config %w", err)
	}
	err = RegisterPgxPool(pool)
	if err != nil {
		return fmt.Errorf("error setting pgx pool %w", err)
	}

	authService := auth_service.NewAuthService(pool)
	err = RegisterAuthService(*authService)
	if err != nil {
		return fmt.Errorf("error setting auth service %w", err)
	}

	authHandlersHTTP := auth_http.NewAuthHandlersHTTP(authService)
	err = RegisterAuthHandlers(authHandlersHTTP)
	if err != nil {
		return fmt.Errorf("error setting auth handlers %w", err)
	}

	pwdService := pwd_services.NewPwdService(pool)
	err = RegisterPwdService(*pwdService)
	if err != nil {
		return fmt.Errorf("error setting pwd service %w", err)
	}

	pwdHandlers := pwd_http.NewPwdHandlersHTTP(pwdService)
	err = RegisterPwdHandlers(pwdHandlers)
	if err != nil {
		return fmt.Errorf("error setting pwd handlers %w", err)
	}

	return nil
}
