package carddto

import (
	"encoding/json"
	"fmt"
	"gophKeeper/internal/helpers"
	"io"
	"net/http"
)

type SaveCardDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Number      string `json:"number"`
	PinCode     string `json:"pincode"`
	CVV         string `json:"cvv"`
	Expire      string `json:"expire"`
	UserID      int    `json:"user_id"`
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

	// Извлекаем userID из контекста
	userID, err := helpers.GetUserIDFromContext(r.Context())
	if err != nil {
		return SaveCardDTO{}, fmt.Errorf("error GetUserIDFromContext: %w", err)
	}

	saveCardDTO.UserID = userID

	return saveCardDTO, nil
}
