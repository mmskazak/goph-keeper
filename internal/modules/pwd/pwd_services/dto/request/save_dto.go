package request

import (
	"encoding/json"
	"fmt"
	"gophKeeper/internal/helpers"
	"gophKeeper/internal/modules/pwd/pwd_services/value_obj"
	"io"
	"net/http"
)

type SavePwdDTO struct {
	UserID      int                   `json:"user_id"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	Credentials value_obj.Credentials `json:"credentials"`
}

func SavePwdDTOFromHTTP(r *http.Request) (SavePwdDTO, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return SavePwdDTO{}, fmt.Errorf("reading body registration: %w", err)
	}
	var savePwdDTO SavePwdDTO
	err = json.Unmarshal(data, &savePwdDTO)
	if err != nil {
		return SavePwdDTO{}, fmt.Errorf("unmarshalling body registration: %w", err)
	}

	// Извлекаем userID из контекста
	userID, err := helpers.getUserIDFromContext(r.Context())
	if err != nil {
		return SavePwdDTO{}, fmt.Errorf("error getUserIDFromContext: %w", err)
	}

	savePwdDTO.UserID = userID

	return savePwdDTO, nil
}
