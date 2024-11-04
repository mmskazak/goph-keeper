package filedto

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"goph-keeper/internal/helpers"
	"net/http"
)

type DeleteFileDTO struct {
	FileID string `json:"file_id"`
	UserID int    `json:"user_id"`
}

func DeleteFileDTOFromHTTP(r *http.Request) (DeleteFileDTO, error) {
	// Извлекаем file_id из пути запроса
	fileID := chi.URLParam(r, "file_id")
	if fileID == "" {
		return DeleteFileDTO{}, errors.New("file_id not found in the request path")
	}

	// Извлекаем userID из контекста
	userID, err := helpers.GetUserIDFromContext(r.Context())
	if err != nil {
		return DeleteFileDTO{}, fmt.Errorf("error GetUserIDFromContext DeleteFileDTOFromHTTP: %w", err)
	}

	var deletePwdDTO DeleteFileDTO
	deletePwdDTO.UserID = userID
	deletePwdDTO.FileID = fileID
	return deletePwdDTO, nil
}
