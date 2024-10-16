package card_dto

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type DeleteCardDTO struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      string `json:"user_id"`
	Number      string `json:"number"`
	PinCode     string `json:"pinCode"`
	Cvv         string `json:"cvv"`
}

func DeleteCardDTOFromHTTP(r *http.Request) (DeleteCardDTO, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return DeleteCardDTO{}, fmt.Errorf("reading body registration: %w", err)
	}
	var deleteTextDTO DeleteCardDTO
	err = json.Unmarshal(data, &deleteTextDTO)
	if err != nil {
		return DeleteCardDTO{}, fmt.Errorf("unmarshalling body registration: %w", err)
	}
	return deleteTextDTO, nil
}
