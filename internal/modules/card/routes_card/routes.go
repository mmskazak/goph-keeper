package routes_card

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"gophKeeper/internal/config"
	"gophKeeper/internal/modules/card/card_http"
	"gophKeeper/internal/modules/card/card_services"
	"net/http"
)

func RegistrationRoutesCard(
	r chi.Router,
	pool *pgxpool.Pool,
	cfg *config.Config,
) chi.Router {
	// Сохранить кредитную карточку
	r.Post("/card/save", func(w http.ResponseWriter, req *http.Request) {
		getCardHandlers(pool).SaveCard(w, req)
	})

	// Получить кредитную карточку
	r.Get("/card/get/{card_id}", func(w http.ResponseWriter, req *http.Request) {
		getCardHandlers(pool).GetCard(w, req)
	})

	// Обновить кредитную карточку
	r.Put("/card/update", func(w http.ResponseWriter, req *http.Request) {
		getCardHandlers(pool).UpdateCard(w, req)
	})

	// Удалить кредитную карточку
	r.Delete("/card/delete", func(w http.ResponseWriter, req *http.Request) {
		getCardHandlers(pool).DeleteCard(w, req)
	})

	// Получить все кредитные карточки
	r.Post("/card/all", func(w http.ResponseWriter, req *http.Request) {
		getCardHandlers(pool).GetAllCards(w, req)
	})

	return r
}

func getCardHandlers(pool *pgxpool.Pool) *card_http.CardHandlers {
	cardService := card_services.NewCardService(pool)
	cardHandlers := card_http.NewCardHandlersHTTP(cardService)
	return &cardHandlers
}
