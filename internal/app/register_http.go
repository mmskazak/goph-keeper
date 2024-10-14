package app

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"gophKeeper/internal/config"
	"gophKeeper/internal/modules/auth/auth_http"
	"gophKeeper/internal/modules/auth/auth_middleware"
	"gophKeeper/internal/modules/auth/auth_services/auth_service"
	"gophKeeper/internal/modules/pwd/pwd_http"
	"gophKeeper/internal/modules/pwd/pwd_services"
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

		//Пароли
		r.Post("/pwd/save", func(w http.ResponseWriter, req *http.Request) {
			getPwdHandlers(pool).SavePassword(w, req)
		})
		r.Post("/pwd/all", func(w http.ResponseWriter, req *http.Request) {
			getPwdHandlers(pool).GetAllPasswords(w, req)
		})
		r.Delete("/pwd/delete", func(w http.ResponseWriter, req *http.Request) {
			getPwdHandlers(pool).DeletePassword(w, req)
		})
		r.Get("/pwd/get", func(w http.ResponseWriter, req *http.Request) {
			getPwdHandlers(pool).GetPassword(w, req)
		})

	})

	return r
}

func getAuthHandlers(pool *pgxpool.Pool, secretKey string) *auth_http.AuthHandlers {
	authService := auth_service.NewAuthService(pool)
	authHandlers := auth_http.NewAuthHandlersHTTP(authService, secretKey)
	return &authHandlers
}

func getPwdHandlers(pool *pgxpool.Pool) *pwd_http.PwdHandlers {
	pwdService := pwd_services.NewPwdService(pool)
	pwdHandlers := pwd_http.NewPwdHandlersHTTP(pwdService)
	return &pwdHandlers
}
