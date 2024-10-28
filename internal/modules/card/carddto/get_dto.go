package carddto

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"gophKeeper/internal/helpers"

	"github.com/go-chi/chi/v5"
)

type GetCardDTO struct {
	CardID int `json:"card_id"` // PwdID карточки в базе данных
	UserID int `json:"user_id"` // Идентификатор пользователя
}

func GetCardDTOFromHTTP(r *http.Request) (GetCardDTO, error) {
	// Извлекаем cardID из пути запроса (пример: card_id)
	cardIDStr := chi.URLParam(r, "card_id")
	if cardIDStr == "" {
		return GetCardDTO{}, errors.New("card_id not found in the request path")
	}

	// Преобразуем cardID в int
	cardID, err := strconv.Atoi(cardIDStr)
	if err != nil {
		return GetCardDTO{}, fmt.Errorf("invalid card_id: %w", err)
	}

	// Извлекаем userID из контекста
	userID, err := helpers.GetUserIDFromContext(r.Context())
	if err != nil {
		return GetCardDTO{}, fmt.Errorf("error GetCardDTOFromHTTP GetUserIDFromContext : %w", err)
	}

	// Формируем DTO
	var getCardDTO GetCardDTO
	getCardDTO.CardID = cardID
	getCardDTO.UserID = userID

	return getCardDTO, nil // Возвращаем DTO и nil (если нет ошибки)
}
