package file_dto

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type UpdateTextDTO struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PathToFile  string `json:"path_to_file"`
}

func UpdateTextDTOFromHTTP(r *http.Request) (UpdateTextDTO, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return UpdateTextDTO{}, fmt.Errorf("reading body registration: %w", err)
	}
	var allTextDTO UpdateTextDTO
	err = json.Unmarshal(data, &allTextDTO)
	if err != nil {
		return UpdateTextDTO{}, fmt.Errorf("unmarshalling body registration: %w", err)
	}
	return allTextDTO, nil
}
