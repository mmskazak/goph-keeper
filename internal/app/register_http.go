package app

import (
	"context"
	"github.com/go-chi/chi/v5"
	"gophKeeper/internal/dig"
	"gophKeeper/internal/modules/auth/auth_http"
	"gophKeeper/internal/modules/auth/auth_middleware"
	"net/http"
)

func registrationHandlersHTTP(
	_ context.Context,
	r *chi.Mux,
) *chi.Mux {
	r.Group(func(r chi.Router) {
		r.Post("/registration", func(w http.ResponseWriter, r *http.Request) {
			getAuthHandlers(w).Registration(w, r)
		})
		r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
			getAuthHandlers(w).Login(w, r)
		})
	})

	r.Group(func(r chi.Router) {
		r.Use(auth_middleware.Authentication)

		r.Get("/logout", func(w http.ResponseWriter, r *http.Request) {
			getAuthHandlers(w).Logout(w, r)
		})
	})

	return r
}

func getAuthHandlers(w http.ResponseWriter) *auth_http.AuthHandlers {
	authHandlers, err := dig.GetAuthHandlers()
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}
	return authHandlers
}
