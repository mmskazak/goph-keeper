package text_dto

import (
	"fmt"
	"gophKeeper/internal/helpers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type GetTextDTO struct {
	TextID string `json:"text_id"`
	UserID int    `json:"user_id"`
}

func GetTextDTOFromHTTP(r *http.Request) (GetTextDTO, error) {
	// Извлекаем текстовый ID из пути запроса (пример: text_id)
	textID := chi.URLParam(r, "text_id")
	if textID == "" {
		return GetTextDTO{}, fmt.Errorf("text_id not found in the request path")
	}

	// Извлекаем userID из контекста
	userID, err := helpers.GetUserIDFromContext(r.Context())
	if err != nil {
		return GetTextDTO{}, fmt.Errorf("error GetUserIDFromContext: %w", err)
	}

	// Формируем DTO
	var getTextDTO GetTextDTO
	// Устанавливаем полученные значения в DTO
	getTextDTO.TextID = textID
	getTextDTO.UserID = userID

	return getTextDTO, nil // Возвращаем DTO и nil (если нет ошибки)
}
