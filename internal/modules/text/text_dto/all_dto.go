package text_dto

import (
	"fmt"
	"gophKeeper/internal/helpers"
	"net/http"
)

type AllTextDTO struct {
	UserID int `json:"user_id"`
}

func AllTextDTOFromHTTP(r *http.Request) (AllTextDTO, error) {
	userID, err := helpers.GetUserIDFromContext(r.Context())
	if err != nil {
		return AllTextDTO{}, fmt.Errorf("error GetUserIDFromContext: %w", err)
	}
	return AllTextDTO{
		UserID: userID,
	}, nil
}
