package pwd_dto

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type SavePwdDTO struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Credentials string `json:"credentials"`
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
	return savePwdDTO, nil
}
