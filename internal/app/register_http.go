package app

import (
	"context"
	"github.com/go-chi/chi/v5"
	"gophKeeper/internal/logger"
	"net/http"
)

func registrationHandlersHTTP(
	_ context.Context,
	router *chi.Mux,
) *chi.Mux {
	router.Post("/registration", func(w http.ResponseWriter, r *http.Request) {
		authHandlers, err := GetAuthHandlersHTTP()
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
		}
		logger.Log.Infoln("Handling registration request")
		authHandlers.Registration(w, r)
	})

	router.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		authHandlers, err := GetAuthHandlersHTTP()
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
		}
		authHandlers.Login(w, r)
	})

	router.Get("/logout", func(w http.ResponseWriter, r *http.Request) {
		authHandlers, err := GetAuthHandlersHTTP()
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
		}
		authHandlers.Logout(w, r)
	})
	return router
}
