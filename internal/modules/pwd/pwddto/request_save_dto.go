package pwddto

import (
	"encoding/json"
	"errors"
	"fmt"
	"goph-keeper/internal/helpers"
	"goph-keeper/internal/modules/pwd/valueobj"
	"io"
	"net/http"
)

type SavePwdDTO struct {
	Credentials valueobj.Credentials `json:"credentials"`
	Title       string               `json:"title"`
	UserID      int                  `json:"user_id"`
}

func SavePwdDTOFromHTTP(r *http.Request) (SavePwdDTO, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return SavePwdDTO{}, fmt.Errorf("reading body for SavePwdDTOFromHTTP: %w", err)
	}
	if len(data) == 0 {
		return SavePwdDTO{}, errors.New("request body is empty")
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
