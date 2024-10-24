package carddto

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GetAllCardsDTO struct {
	UserID int `json:"user_id"` // Идентификатор пользователя
}

func GetAllCardsDTOFromHTTP(r *http.Request) (GetAllCardsDTO, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return GetAllCardsDTO{}, fmt.Errorf("reading body for get cards dto: %w", err)
	}
	var getAllCardsDTO GetAllCardsDTO
	err = json.Unmarshal(data, &getAllCardsDTO)
	if err != nil {
		return GetAllCardsDTO{}, fmt.Errorf("unmarshalling body for get cards dto: %w", err)
	}
	return getAllCardsDTO, nil
}
