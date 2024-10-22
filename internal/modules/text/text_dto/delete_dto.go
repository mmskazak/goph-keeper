package text_dto

import (
	"encoding/json"
	"fmt"
	"gophKeeper/internal/helpers"
	"io"
	"net/http"
)

type DeleteTextDTO struct {
	TextID string `json:"text_id"`
	UserID int    `json:"user_id"`
}

func DeleteTextDTOFromHTTP(r *http.Request) (DeleteTextDTO, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return DeleteTextDTO{}, fmt.Errorf("reading body registration: %w", err)
	}
	var deleteTextDTO DeleteTextDTO
	err = json.Unmarshal(data, &deleteTextDTO)
	if err != nil {
		return DeleteTextDTO{}, fmt.Errorf("unmarshalling body registration: %w", err)
	}

	userID, err := helpers.GetUserIDFromContext(r.Context())
	if err != nil {
		return DeleteTextDTO{}, fmt.Errorf("error GetUserIDFromContext: %w", err)
	}

	deleteTextDTO.UserID = userID
	return deleteTextDTO, nil
}
