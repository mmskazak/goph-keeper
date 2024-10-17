package request

import (
	"fmt"
	"net/http"
)

type AllPwdDTO struct {
	UserID int `json:"user_id"`
}

func AllPwdDTOFromHTTP(r *http.Request) (AllPwdDTO, error) {
	// Извлекаем userID из контекста
	userID, err := getUserIDFromContext(r.Context())
	if err != nil {
		return AllPwdDTO{}, fmt.Errorf("error getUserIDFromContext: %w", err)
	}
	return AllPwdDTO{
		UserID: userID,
	}, nil
}