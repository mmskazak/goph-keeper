package request

import (
	"encoding/json"
	"fmt"
	"gophKeeper/internal/helpers"
	"io"
	"net/http"
)

type SaveFileDTO struct {
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
	UserID   int    `json:"user_id"`
}

func SaveFileDTOFromHTTP(r *http.Request) (SaveFileDTO, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return SaveFileDTO{}, fmt.Errorf("reading body registration: %w", err)
	}
	var savePwdDTO SaveFileDTO
	err = json.Unmarshal(data, &savePwdDTO)
	if err != nil {
		return SaveFileDTO{}, fmt.Errorf("unmarshalling body registration: %w", err)
	}

	// Извлекаем userID из контекста
	userID, err := helpers.GetUserIDFromContext(r.Context())
	if err != nil {
		return SaveFileDTO{}, fmt.Errorf("error GetUserIDFromContext: %w", err)
	}

	savePwdDTO.UserID = userID

	return savePwdDTO, nil
}
