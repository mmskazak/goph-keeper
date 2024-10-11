package app

import (
	"context"
	"github.com/go-chi/chi/v5"
	"gophKeeper/internal/logger"
	"gophKeeper/internal/modules/auth/http"
	"gophKeeper/internal/service_locator"
	"net/http"
)

func registrationHandlersHTTP(
	_ context.Context,
	router *chi.Mux,
) *chi.Mux {
	router.Post("/registration", func(w http.ResponseWriter, r *http.Request) {
		logger.Log.Infoln("Handling registration request")
		getAuthHandlers(w).Registration(w, r)
	})

	router.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		getAuthHandlers(w).Login(w, r)
	})

	router.Get("/logout", func(w http.ResponseWriter, r *http.Request) {
		getAuthHandlers(w).Logout(w, r)
	})
	return router
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
