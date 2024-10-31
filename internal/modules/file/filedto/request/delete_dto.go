package request

import (
	"encoding/json"
	"fmt"
	"goph-keeper/internal/helpers"
	"io"
	"net/http"
)

type DeleteFileDTO struct {
	FileID string `json:"file_id"`
	UserID int    `json:"user_id"`
}

func DeleteFileDTOFromHTTP(r *http.Request) (DeleteFileDTO, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return DeleteFileDTO{}, fmt.Errorf("reading body for delete file dto: %w", err)
	}
	var deletePwdDTO DeleteFileDTO
	err = json.Unmarshal(data, &deletePwdDTO)
	if err != nil {
		return DeleteFileDTO{}, fmt.Errorf("unmarshalling body registration: %w", err)
	}

	// Извлекаем userID из контекста
	userID, err := helpers.GetUserIDFromContext(r.Context())
	if err != nil {
		return DeleteFileDTO{}, fmt.Errorf("error GetUserIDFromContext DeleteFileDTOFromHTTP: %w", err)
	}

	deletePwdDTO.UserID = userID
	return deletePwdDTO, nil
}
