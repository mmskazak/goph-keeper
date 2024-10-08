package app

import (
	"github.com/go-chi/chi/v5"
	"gophKeeper/internal/config"
)

// registrationMiddleware регистрация посредников.
func registrationMiddleware(router *chi.Mux, _ *config.Config) *chi.Mux {
	//TODO: добавить посредников
	return router
}
