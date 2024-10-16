package card_dto

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GetCardDTO struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      string `json:"user_id"`
	Number      string `json:"number"`
	PinCode     string `json:"pinCode"`
	Cvv         string `json:"cvv"`
}

func GetCardDTOFromHTTP(r *http.Request) (GetCardDTO, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return GetCardDTO{}, fmt.Errorf("reading body registration: %w", err)
	}
	var getTextDTO GetCardDTO
	err = json.Unmarshal(data, &getTextDTO)
	if err != nil {
		return GetCardDTO{}, fmt.Errorf("unmarshalling body registration: %w", err)
	}
	return getTextDTO, nil
}
