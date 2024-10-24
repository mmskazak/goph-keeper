package request

import (
	"encoding/json"
	"fmt"
	"gophKeeper/internal/helpers"
	"gophKeeper/internal/modules/pwd/valueobj"
	"io"
	"net/http"
)

type UpdatePwdDTO struct {
	ID          string               `json:"pwd_id"`
	UserID      int                  `json:"user_id"`
	Title       string               `json:"title"`
	Description string               `json:"description"`
	Credentials valueobj.Credentials `json:"credentials"`
}

func UpdatePwdDTOFromHTTP(r *http.Request) (UpdatePwdDTO, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return UpdatePwdDTO{}, fmt.Errorf("reading body registration: %w", err)
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
