package textdto

import (
	"encoding/json"
	"fmt"
	"goph-keeper/internal/helpers"
	"io"
	"net/http"
)

type SaveTextDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	TextContent string `json:"text_content"`
	UserID      int    `json:"user_id"`
}

func SaveTextDTOFromHTTP(r *http.Request) (SaveTextDTO, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return SaveTextDTO{}, fmt.Errorf("reading body for SaveTextDTOFromHTTP: %w", err)
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
