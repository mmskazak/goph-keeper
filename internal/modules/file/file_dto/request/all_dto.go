package request

import (
	"fmt"
	"gophKeeper/internal/helpers"
	"net/http"
)

type AllFilesDTO struct {
	UserID int `json:"user_id"`
}

func AllFileDTOFromHTTP(r *http.Request) (AllFilesDTO, error) {
	// Извлекаем userID из контекста
	userID, err := helpers.GetUserIDFromContext(r.Context())
	if err != nil {
		return AllFilesDTO{}, fmt.Errorf("error GetUserIDFromContext: %w", err)
	}
	return AllFilesDTO{
		UserID: userID,
	}, nil
}
