package app

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"gophKeeper/internal/config"
	"gophKeeper/internal/modules/auth/auth_http"
	"gophKeeper/internal/modules/auth/auth_middleware"
	"gophKeeper/internal/modules/auth/auth_services/auth_service"
	"gophKeeper/internal/modules/file/routes_file"
	"gophKeeper/internal/modules/pwd/routes_pwd"
	"net/http"
)

func registrationHandlersHTTP(
	_ context.Context,
	r *chi.Mux,
	cfg *config.Config,
	pool *pgxpool.Pool,
) *chi.Mux {
	r.Group(func(r chi.Router) {
		r.Post("/registration", func(w http.ResponseWriter, req *http.Request) {
			getAuthHandlers(pool, cfg.SecretKey).Registration(w, req)
		})
		r.Post("/login", func(w http.ResponseWriter, req *http.Request) {
			getAuthHandlers(pool, cfg.SecretKey).Login(w, req)
		})
	})

	r.Group(func(r chi.Router) {
		r.Use(func(next http.Handler) http.Handler {
			return auth_middleware.Authentication(next, cfg.SecretKey)
		})

		r.Get("/logout", func(w http.ResponseWriter, r *http.Request) {
			getAuthHandlers(pool, cfg.SecretKey).Logout(w, r)
		})

		r = routes_pwd.RegistrationRoutesPwd(r, pool, cfg)
		r = routes_file.RegistrationRoutesFile(r, pool, cfg)
	})

	return r
}

func getAuthHandlers(pool *pgxpool.Pool, secretKey string) *auth_http.AuthHandlers {
	authService := auth_service.NewAuthService(pool)
	authHandlers := auth_http.NewAuthHandlersHTTP(authService, secretKey)
	return &authHandlers
}
