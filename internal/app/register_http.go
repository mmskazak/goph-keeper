package app

import (
	"context"
	"github.com/go-chi/chi/v5"
	"gophKeeper/internal/modules/auth/http"
	"gophKeeper/internal/modules/auth/middleware"
	"gophKeeper/internal/service_locator"
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
		r.Use(middleware.Authentication)

		r.Get("/logout", func(w http.ResponseWriter, r *http.Request) {
			getAuthHandlers(w).Logout(w, r)
		})
	})

	return r
}

func getAuthHandlers(w http.ResponseWriter) *auth_http.AuthHandlers {
	sl := service_locator.InitServiceLocator()
	ah := sl.Get("auth_handlers_http")
	authHandlers, ok := ah.(auth_http.AuthHandlers)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return nil
	}

	return &authHandlers
}
