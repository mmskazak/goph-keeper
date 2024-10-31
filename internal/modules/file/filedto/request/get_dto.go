package request

import (
	"errors"
	"fmt"
	"goph-keeper/internal/helpers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type GetFileDTO struct {
	FileID string `json:"file_id"` // PwdID пароля в базе данных
	UserID int    `json:"user_id"` // Идентификатор пользователя
}

func GetFileDTOFromHTTP(r *http.Request) (GetFileDTO, error) {
	// Извлекаем file_id из пути запроса
	fileID := chi.URLParam(r, "file_id")
	if fileID == "" {
		return GetFileDTO{}, errors.New("file_id not found in the request path")
	}

	// Извлекаем userID из контекста
	userID, err := helpers.GetUserIDFromContext(r.Context())
	if err != nil {
		return GetFileDTO{}, fmt.Errorf("error GetUserIDFromContext GetFileDTOFromHTTP: %w", err)
	}

	// Формируем DTO
	getFileDTO := GetFileDTO{
		FileID: fileID,
		UserID: userID,
	}

	return getFileDTO, nil // Возвращаем DTO и nil (если нет ошибки)
}
