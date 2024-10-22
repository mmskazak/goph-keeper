package text_dto

import (
	"encoding/json"
	"fmt"
	"gophKeeper/internal/helpers"
	"io"
	"net/http"
)

type SaveTextDTO struct {
	UserID      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	TextContent string `json:"text_content"`
}

func SaveTextDTOFromHTTP(r *http.Request) (SaveTextDTO, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return SaveTextDTO{}, fmt.Errorf("reading body registration: %w", err)
	}
	var saveTextDTO SaveTextDTO
	err = json.Unmarshal(data, &saveTextDTO)
	if err != nil {
		return SaveTextDTO{}, fmt.Errorf("unmarshalling body registration: %w", err)
	}

	// Извлекаем userID из контекста
	userID, err := helpers.GetUserIDFromContext(r.Context())
	if err != nil {
		return SaveTextDTO{}, fmt.Errorf("error GetUserIDFromContext: %w", err)
	}

	saveTextDTO.UserID = userID

	return saveTextDTO, nil
}
