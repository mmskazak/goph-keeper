package carddto

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type SaveCardDTO struct {
	UserID      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Number      string `json:"number"`
	PinCode     string `json:"pincode"`
	CVV         string `json:"cvv"`
	Expire      string `json:"expire"`
}

func SaveCardDTOFromHTTP(r *http.Request) (SaveCardDTO, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return SaveCardDTO{}, fmt.Errorf("reading body for save card dto: %w", err)
	}
	var saveCardDTO SaveCardDTO
	err = json.Unmarshal(data, &saveCardDTO)
	if err != nil {
		return SaveCardDTO{}, fmt.Errorf("unmarshalling body for save card dto: %w", err)
	}
	return saveCardDTO, nil
}
