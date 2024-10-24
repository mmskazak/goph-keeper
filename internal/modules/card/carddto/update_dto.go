package carddto

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type UpdateCardDTO struct {
	Title       string `json:"title"`       // Новое название
	Description string `json:"description"` // Новое описание
	Number      string `json:"number"`      // Новый номер
	PinCode     string `json:"pincode"`     // Новый ПИН-код
	CVV         string `json:"cvv"`         // Новый CVV
	Expire      string `json:"expire"`      // Новая дата истечения
	CardID      int    `json:"card_id"`     // ID карточки
}

func UpdateCardDTOFromHTTP(r *http.Request) (UpdateCardDTO, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return UpdateCardDTO{}, fmt.Errorf("reading body for update card dto: %w", err)
	}
	var updateCardDTO UpdateCardDTO
	err = json.Unmarshal(data, &updateCardDTO)
	if err != nil {
		return UpdateCardDTO{}, fmt.Errorf("unmarshalling body for update card dto: %w", err)
	}
	return updateCardDTO, nil
}
