package app

import (
	"GophKeeper/internal/modules/auth/contract"
	"context"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func registrationHTTP(
	_ context.Context,
	router *chi.Mux,
	serviceHttp contract.IAuth,
) *chi.Mux {
	router.Post("/registration", func(w http.ResponseWriter, r *http.Request) {
		serviceHttp.Registration(w, r)
	})

	router.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		serviceHttp.Login(w, r)
	})

	router.Get("/logout", func(w http.ResponseWriter, r *http.Request) {
		serviceHttp.Logout(w, r)
	})
	return router
}
