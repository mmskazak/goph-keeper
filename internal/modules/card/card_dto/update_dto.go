package card_dto

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type UpdateCardDTO struct {
	CardID      int    `json:"card_id"`     // ID карточки
	Title       string `json:"title"`       // Новое название
	Description string `json:"description"` // Новое описание
	Number      string `json:"number"`      // Новый номер
	PinCode     string `json:"pincode"`     // Новый ПИН-код
	CVV         string `json:"cvv"`         // Новый CVV
	Expire      string `json:"expire"`      // Новая дата истечения
}

func UpdateCardDTOFromHTTP(r *http.Request) (UpdateCardDTO, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return UpdateCardDTO{}, fmt.Errorf("reading body: %w", err)
	}
	var updateCardDTO UpdateCardDTO
	err = json.Unmarshal(data, &updateCardDTO)
	if err != nil {
		return UpdateCardDTO{}, fmt.Errorf("unmarshalling body: %w", err)
	}
	return updateCardDTO, nil
}
