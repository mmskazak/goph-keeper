package routestext

import (
	"goph-keeper/internal/config"
	"goph-keeper/internal/modules/text/texthttp"
	"goph-keeper/internal/modules/text/textservices"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegistrationRoutesText(
	r chi.Router,
	pool *pgxpool.Pool,
	cfg *config.Config,
) {
	// Сохранить текст
	r.Post("/text/save", func(w http.ResponseWriter, req *http.Request) {
		getTextHandlers(pool, cfg.EncryptionKey).SaveText(w, req)
	})
	// Получить все тексты
	r.Post("/text/all", func(w http.ResponseWriter, req *http.Request) {
		getTextHandlers(pool, cfg.EncryptionKey).GetAllTexts(w, req)
	})
	// Удалить текст
	r.Post("/text/delete", func(w http.ResponseWriter, req *http.Request) {
		getTextHandlers(pool, cfg.EncryptionKey).DeleteText(w, req)
	})
	// Получить конкретный текст
	r.Get("/text/get/{text_id}", func(w http.ResponseWriter, req *http.Request) {
		getTextHandlers(pool, cfg.EncryptionKey).GetText(w, req)
	})
	// Обновить конкретный текст
	r.Post("/text/update", func(w http.ResponseWriter, req *http.Request) {
		getTextHandlers(pool, cfg.EncryptionKey).UpdateText(w, req)
	})
}

func getTextHandlers(pool *pgxpool.Pool, enKey [32]byte) *texthttp.TextHandlers {
	textService := textservices.NewTextService(pool, enKey)
	textHandlers := texthttp.NewTextHandlersHTTP(textService)
	return &textHandlers
}
