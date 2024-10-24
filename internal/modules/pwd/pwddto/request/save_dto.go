package request

import (
	"encoding/json"
	"fmt"
	"gophKeeper/internal/helpers"
	"gophKeeper/internal/modules/pwd/valueobj"
	"io"
	"net/http"
)

type SavePwdDTO struct {
	Title       string               `json:"title"`
	Description string               `json:"description"`
	UserID      int                  `json:"user_id"`
	Credentials valueobj.Credentials `json:"credentials"`
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
	userID, err := helpers.GetUserIDFromContext(r.Context())
	if err != nil {
		return SavePwdDTO{}, fmt.Errorf("error GetUserIDFromContext: %w", err)
	}

	savePwdDTO.UserID = userID

	return savePwdDTO, nil
}
