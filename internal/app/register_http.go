package app

import (
	"context"
	"goph-keeper/internal/config"
	"goph-keeper/internal/modules/auth/authhttp"
	"goph-keeper/internal/modules/auth/authmiddleware"
	"goph-keeper/internal/modules/auth/authservices/authservice"
	"goph-keeper/internal/modules/file/routesfile"
	"goph-keeper/internal/modules/pwd/routespwd"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
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
			return authmiddleware.Authentication(next, cfg.SecretKey)
		})

		routespwd.RegistrationRoutesPwd(r, pool, cfg)
		routesfile.RegistrationRoutesFile(r, pool, cfg)
	})

	return r
}

func getAuthHandlers(pool *pgxpool.Pool, secretKey string) *authhttp.AuthHandlers {
	authService := authservice.NewAuthService(pool)
	authHandlers := authhttp.NewAuthHandlersHTTP(authService, secretKey)
	return &authHandlers
}
