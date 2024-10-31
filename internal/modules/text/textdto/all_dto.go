package textdto

import (
	"fmt"
	"goph-keeper/internal/helpers"
	"net/http"
)

type AllTextDTO struct {
	UserID int `json:"user_id"`
}

func AllTextDTOFromHTTP(r *http.Request) (AllTextDTO, error) {
	userID, err := helpers.GetUserIDFromContext(r.Context())
	if err != nil {
		return AllTextDTO{}, fmt.Errorf("error AllTextDTOFromHTTP GetUserIDFromContext: %w", err)
	}
	return AllTextDTO{
		UserID: userID,
	}, nil
}
