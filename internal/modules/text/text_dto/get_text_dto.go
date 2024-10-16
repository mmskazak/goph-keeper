package text_dto

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GetTextDTO struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Text        string `json:"text"`
}

func GetTextDTOFromHTTP(r *http.Request) (GetTextDTO, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return GetTextDTO{}, fmt.Errorf("reading body registration: %w", err)
	}
	var getTextDTO GetTextDTO
	err = json.Unmarshal(data, &getTextDTO)
	if err != nil {
		return GetTextDTO{}, fmt.Errorf("unmarshalling body registration: %w", err)
	}
	return getTextDTO, nil
}
