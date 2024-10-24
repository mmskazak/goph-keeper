package routespwd

import (
	"gophKeeper/internal/config"
	"gophKeeper/internal/modules/pwd/pwdhttp"
	"gophKeeper/internal/modules/pwd/pwdservices"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegistrationRoutesPwd(
	r chi.Router,
	pool *pgxpool.Pool,
	cfg *config.Config,
) {
	// Сохранить пароль
	r.Post("/pwd/save", func(w http.ResponseWriter, req *http.Request) {
		getPwdHandlers(pool, cfg.EncryptionKey).SavePassword(w, req)
	})
	// Получить все пароли
	r.Post("/pwd/all", func(w http.ResponseWriter, req *http.Request) {
		getPwdHandlers(pool, cfg.EncryptionKey).GetAllPasswords(w, req)
	})
	// Удалить пароль
	r.Delete("/pwd/delete", func(w http.ResponseWriter, req *http.Request) {
		getPwdHandlers(pool, cfg.EncryptionKey).DeletePassword(w, req)
	})
	// Получить конкретный пароль
	r.Post("/pwd/get", func(w http.ResponseWriter, req *http.Request) {
		getPwdHandlers(pool, cfg.EncryptionKey).GetPassword(w, req)
	})

	// Получить конкретный пароль
	r.Post("/pwd/update", func(w http.ResponseWriter, req *http.Request) {
		getPwdHandlers(pool, cfg.EncryptionKey).UpdatePassword(w, req)
	})
}

func getPwdHandlers(pool *pgxpool.Pool, enKey [32]byte) *pwdhttp.PwdHandlers {
	pwdService := pwdservices.NewPwdService(pool, enKey)
	pwdHandlers := pwdhttp.NewPwdHandlersHTTP(pwdService)
	return &pwdHandlers
}
