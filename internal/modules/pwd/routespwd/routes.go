package routespwd

import (
	"goph-keeper/internal/config"
	"goph-keeper/internal/modules/pwd/pwdhttp"
	"goph-keeper/internal/modules/pwd/pwdservices"
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
	// Получить конкретный пароль
	r.Post("/pwd/update", func(w http.ResponseWriter, req *http.Request) {
		getPwdHandlers(pool, cfg.EncryptionKey).UpdatePassword(w, req)
	})
	// Получить все пароли
	r.Get("/pwd/all", func(w http.ResponseWriter, req *http.Request) {
		getPwdHandlers(pool, cfg.EncryptionKey).GetAllPasswords(w, req)
	})
	// Удалить пароль
	r.Get("/pwd/delete/{pwd_id}", func(w http.ResponseWriter, req *http.Request) {
		getPwdHandlers(pool, cfg.EncryptionKey).DeletePassword(w, req)
	})
	// Получить конкретный пароль
	r.Get("/pwd/get/{pwd_id}", func(w http.ResponseWriter, req *http.Request) {
		getPwdHandlers(pool, cfg.EncryptionKey).GetPassword(w, req)
	})
}

func getPwdHandlers(pool *pgxpool.Pool, enKey [32]byte) *pwdhttp.PwdHandlers {
	pwdService := pwdservices.NewPwdService(pool, enKey)
	pwdHandlers := pwdhttp.NewPwdHandlersHTTP(pwdService)
	return &pwdHandlers
}
