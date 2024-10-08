package app

import (
	"context"
	"github.com/go-chi/chi/v5"
	servive_http "gophKeeper/internal/modules/auth/http"

	"net/http"
)

func registrationHandlersHTTP(
	_ context.Context,
	router *chi.Mux,
) *chi.Mux {
	router.Post("/registration", func(w http.ResponseWriter, r *http.Request) {
		authServiceHTTP := servive_http.NewAuthServiceHTTP()
		authServiceHTTP.Registration(w, r)
	})

	router.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		authServiceHTTP := servive_http.NewAuthServiceHTTP()
		authServiceHTTP.Login(w, r)
	})

	router.Get("/logout", func(w http.ResponseWriter, r *http.Request) {
		authServiceHTTP := servive_http.NewAuthServiceHTTP()
		authServiceHTTP.Logout(w, r)
	})
	return router
}
