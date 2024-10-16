package card_dto

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type AllCardDTO struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      string `json:"user_id"`
	Number      string `json:"number"`
	PinCode     string `json:"pinCode"`
	Cvv         string `json:"cvv"`
}

func AllCardsDTOFromHTTP(r *http.Request) (AllCardDTO, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return AllCardDTO{}, fmt.Errorf("reading body registration: %w", err)
	}
	var allTextDTO AllCardDTO
	err = json.Unmarshal(data, &allTextDTO)
	if err != nil {
		return AllCardDTO{}, fmt.Errorf("unmarshalling body registration: %w", err)
	}
	return allTextDTO, nil
}
