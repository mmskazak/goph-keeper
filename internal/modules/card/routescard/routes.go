package routescard

import (
	"gophKeeper/internal/config"
	"gophKeeper/internal/modules/card/cardhttp"
	"gophKeeper/internal/modules/card/cardservices"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegistrationRoutesCard(
	r chi.Router,
	pool *pgxpool.Pool,
	_ *config.Config,
) {
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
}

func getCardHandlers(pool *pgxpool.Pool) *cardhttp.CardHandlers {
	cardService := cardservices.NewCardService(pool)
	cardHandlers := cardhttp.NewCardHandlersHTTP(cardService)
	return &cardHandlers
}
