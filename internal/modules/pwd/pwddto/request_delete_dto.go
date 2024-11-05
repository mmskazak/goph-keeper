package pwddto

import (
	"errors"
	"fmt"
	"goph-keeper/internal/helpers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type DeletePwdDTO struct {
	PwdID  string `json:"pwd_id"`
	UserID int    `json:"user_id"`
}

func DeletePwdDTOFromHTTP(r *http.Request) (DeletePwdDTO, error) {
	// Извлекаем текстовый PwdID из пути запроса (пример: text_id)
	pwdID := chi.URLParam(r, "pwd_id")
	if pwdID == "" {
		return DeletePwdDTO{}, errors.New("pwd_id not found in the request path")
	}

	// Извлекаем userID из контекста
	userID, err := helpers.GetUserIDFromContext(r.Context())
	if err != nil {
		return DeletePwdDTO{}, fmt.Errorf("error DeletePwdDTOFromHTTP GetUserIDFromContext: %w", err)
	}

	var deletePwdDTO DeletePwdDTO
	deletePwdDTO.PwdID = pwdID
	deletePwdDTO.UserID = userID
	return deletePwdDTO, nil
}
