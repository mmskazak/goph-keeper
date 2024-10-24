package carddto

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type DeleteCardDTO struct {
	CardID int `json:"card_id"` // ID карточки для удаления
}

func DeleteCardDTOFromHTTP(r *http.Request) (DeleteCardDTO, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return DeleteCardDTO{}, fmt.Errorf("reading body for delete card dto: %w", err)
	}
	var deleteCardDTO DeleteCardDTO
	err = json.Unmarshal(data, &deleteCardDTO)
	if err != nil {
		return DeleteCardDTO{}, fmt.Errorf("unmarshalling body for delete card dto: %w", err)
	}
	return deleteCardDTO, nil
}
