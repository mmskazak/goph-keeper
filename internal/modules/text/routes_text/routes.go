package routes_text

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"gophKeeper/internal/config"
	"gophKeeper/internal/modules/text/text_http"
	"gophKeeper/internal/modules/text/text_services"
	"net/http"
)

func RegistrationRoutesText(
	r chi.Router,
	pool *pgxpool.Pool,
	cfg *config.Config,
) chi.Router {
	// Сохранить текст
	r.Post("/text/save", func(w http.ResponseWriter, req *http.Request) {
		getTextHandlers(pool, cfg.EncryptionKey).SaveText(w, req)
	})
	// Получить все тексты
	r.Post("/text/all", func(w http.ResponseWriter, req *http.Request) {
		getTextHandlers(pool, cfg.EncryptionKey).GetAllTexts(w, req)
	})
	// Удалить текст
	r.Delete("/text/delete", func(w http.ResponseWriter, req *http.Request) {
		getTextHandlers(pool, cfg.EncryptionKey).DeleteText(w, req)
	})
	// Получить конкретный текст
	r.Post("/text/get", func(w http.ResponseWriter, req *http.Request) {
		getTextHandlers(pool, cfg.EncryptionKey).GetText(w, req)
	})
	// Обновить конкретный текст
	r.Post("/text/update", func(w http.ResponseWriter, req *http.Request) {
		getTextHandlers(pool, cfg.EncryptionKey).UpdateText(w, req)
	})

	return r
}

func getTextHandlers(pool *pgxpool.Pool, enKey [32]byte) *text_http.TextHandlers {
	textService := text_services.NewTextService(pool, enKey)
	textHandlers := text_http.NewTextHandlersHTTP(textService)
	return &textHandlers
}
