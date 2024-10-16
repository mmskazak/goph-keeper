package file_dto

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type DeleteTextDTO struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PathToFile  string `json:"path_to_file"`
}

func DeletePwdDTOFromHTTP(r *http.Request) (DeleteTextDTO, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return DeleteTextDTO{}, fmt.Errorf("reading body registration: %w", err)
	}
	var deleteTextDTO DeleteTextDTO
	err = json.Unmarshal(data, &deleteTextDTO)
	if err != nil {
		return DeleteTextDTO{}, fmt.Errorf("unmarshalling body registration: %w", err)
	}
	return deleteTextDTO, nil
}
