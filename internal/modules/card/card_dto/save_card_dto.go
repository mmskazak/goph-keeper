package card_dto

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type SaveCardDTO struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      string `json:"user_id"`
	Number      string `json:"number"`
	PinCode     string `json:"pinCode"`
	Cvv         string `json:"cvv"`
}

func SaveCardDTOFromHTTP(r *http.Request) (SaveCardDTO, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return SaveCardDTO{}, fmt.Errorf("reading body registration: %w", err)
	}
	var saveCardDTO SaveCardDTO
	err = json.Unmarshal(data, &saveCardDTO)
	if err != nil {
		return SaveCardDTO{}, fmt.Errorf("unmarshalling body registration: %w", err)
	}
	return saveCardDTO, nil
}
