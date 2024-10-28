package pwddto

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"gophKeeper/internal/helpers"
	"net/http"
)

type GetPwdDTO struct {
	PwdID  string `json:"pwd_id"`  // PwdID пароля в базе данных
	UserID int    `json:"user_id"` // Идентификатор пользователя
}

func GetPwdDTOFromHTTP(r *http.Request) (GetPwdDTO, error) {
	// Извлекаем текстовый PwdID из пути запроса (пример: text_id)
	pwdID := chi.URLParam(r, "pwd_id")
	if pwdID == "" {
		return GetPwdDTO{}, errors.New("text_id not found in the request path")
	}

	// Извлекаем userID из контекста
	userID, err := helpers.GetUserIDFromContext(r.Context())
	if err != nil {
		return GetPwdDTO{}, fmt.Errorf("error GetUserIDFromContext: %w", err)
	}

	// Формируем DTO
	var getTextDTO GetPwdDTO
	// Устанавливаем полученные значения в DTO
	getTextDTO.PwdID = pwdID
	getTextDTO.UserID = userID

	return getTextDTO, nil
}
