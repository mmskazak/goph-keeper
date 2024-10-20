package request

import (
	"fmt"
	"net/http"
)

type AllFilesDTO struct {
	UserID int `json:"user_id"`
}

func AllFileDTOFromHTTP(r *http.Request) (AllFilesDTO, error) {
	// Извлекаем userID из контекста
	userID, err := getUserIDFromContext(r.Context())
	if err != nil {
		return AllFilesDTO{}, fmt.Errorf("error getUserIDFromContext: %w", err)
	}
	return AllFilesDTO{
		UserID: userID,
	}, nil
}
