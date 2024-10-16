package file_dto

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type SaveTextDTO struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PathToFile  string `json:"path_to_file"`
}

func SaveTextDTOFromHTTP(r *http.Request) (SaveTextDTO, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return SaveTextDTO{}, fmt.Errorf("reading body registration: %w", err)
	}
	var saveTextDTO SaveTextDTO
	err = json.Unmarshal(data, &saveTextDTO)
	if err != nil {
		return SaveTextDTO{}, fmt.Errorf("unmarshalling body registration: %w", err)
	}
	return saveTextDTO, nil
}
