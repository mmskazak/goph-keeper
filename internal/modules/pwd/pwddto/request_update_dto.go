package pwddto

import (
	"encoding/json"
	"fmt"
	"goph-keeper/internal/helpers"
	"goph-keeper/internal/modules/pwd/valueobj"
	"io"
	"net/http"
)

type UpdatePwdDTO struct {
	PwdID       string               `json:"pwd_id"`
	Title       string               `json:"title"`
	UserID      int                  `json:"user_id"`
	Credentials valueobj.Credentials `json:"credentials"`
}

func UpdatePwdDTOFromHTTP(r *http.Request) (UpdatePwdDTO, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return UpdatePwdDTO{}, fmt.Errorf("reading body for UpdatePwdDTOFromHTTP: %w", err)
	}
	var updatePwdDTO UpdatePwdDTO
	err = json.Unmarshal(data, &updatePwdDTO)
	if err != nil {
		return UpdatePwdDTO{}, fmt.Errorf("unmarshalling body registration: %w", err)
	}

	// Извлекаем userID из контекста
	userID, err := helpers.GetUserIDFromContext(r.Context())
	if err != nil {
		return UpdatePwdDTO{}, fmt.Errorf("error GetUserIDFromContext: %w", err)
	}

	updatePwdDTO.UserID = userID

	return updatePwdDTO, nil
}
