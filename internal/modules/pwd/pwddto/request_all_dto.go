package pwddto

import (
	"fmt"
	"goph-keeper/internal/helpers"
	"net/http"
)

type AllPwdDTO struct {
	UserID int `json:"user_id"`
}

func AllPwdDTOFromHTTP(r *http.Request) (AllPwdDTO, error) {
	// Извлекаем userID из контекста
	userID, err := helpers.GetUserIDFromContext(r.Context())
	if err != nil {
		return AllPwdDTO{}, fmt.Errorf("error GetUserIDFromContext AllPwdDTOFromHTTP: %w", err)
	}
	return AllPwdDTO{
		UserID: userID,
	}, nil
}
