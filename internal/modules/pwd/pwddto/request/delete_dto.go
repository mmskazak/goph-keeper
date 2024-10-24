package request

import (
	"encoding/json"
	"fmt"
	"gophKeeper/internal/helpers"
	"io"
	"net/http"
)

type DeletePwdDTO struct {
	PwdID  string `json:"pwd_id"`
	UserID int    `json:"user_id"`
}

func DeletePwdDTOFromHTTP(r *http.Request) (DeletePwdDTO, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return DeletePwdDTO{}, fmt.Errorf("reading body registration: %w", err)
	}
	var deletePwdDTO DeletePwdDTO
	err = json.Unmarshal(data, &deletePwdDTO)
	if err != nil {
		return DeletePwdDTO{}, fmt.Errorf("unmarshalling body registration: %w", err)
	}

	// Извлекаем userID из контекста
	userID, err := helpers.GetUserIDFromContext(r.Context())
	if err != nil {
		return DeletePwdDTO{}, fmt.Errorf("error GetUserIDFromContext: %w", err)
	}

	deletePwdDTO.UserID = userID
	return deletePwdDTO, nil
}
