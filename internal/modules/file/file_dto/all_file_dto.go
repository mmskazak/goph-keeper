package file_dto

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type AllTextDTO struct {
	UserID string `json:"user_id"`
}

func AllTextsDTOFromHTTP(r *http.Request) (AllTextDTO, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return AllTextDTO{}, fmt.Errorf("reading body registration: %w", err)
	}
	var allTextDTO AllTextDTO
	err = json.Unmarshal(data, &allTextDTO)
	if err != nil {
		return AllTextDTO{}, fmt.Errorf("unmarshalling body registration: %w", err)
	}
	return allTextDTO, nil
}
