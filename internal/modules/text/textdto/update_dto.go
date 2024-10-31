package textdto

import (
	"encoding/json"
	"fmt"
	"goph-keeper/internal/helpers"
	"io"
	"net/http"
)

type UpdateTextDTO struct {
	ID          string `json:"text_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	TextContent string `json:"text_content"`
	UserID      int    `json:"user_id"`
}

func UpdateTextDTOFromHTTP(r *http.Request) (UpdateTextDTO, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return UpdateTextDTO{}, fmt.Errorf("reading body for UpdateTextDTOFromHTTP: %w", err)
	}
	var updateTextDTO UpdateTextDTO
	err = json.Unmarshal(data, &updateTextDTO)
	if err != nil {
		return UpdateTextDTO{}, fmt.Errorf("unmarshalling body registration: %w", err)
	}

	userID, err := helpers.GetUserIDFromContext(r.Context())
	if err != nil {
		return UpdateTextDTO{}, fmt.Errorf("error GetUserIDFromContext: %w", err)
	}

	updateTextDTO.UserID = userID

	return updateTextDTO, nil
}
