package pwd_dto

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GetPwdDTO struct {
	UserID string `json:"user_id"`
	Login  string `json:"login"`
}

func GetPwdDTOFromHTTP(r *http.Request) (GetPwdDTO, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return GetPwdDTO{}, fmt.Errorf("reading body registration: %w", err)
	}
	var getPwdDTO GetPwdDTO
	err = json.Unmarshal(data, &getPwdDTO)
	if err != nil {
		return GetPwdDTO{}, fmt.Errorf("unmarshalling body registration: %w", err)
	}
	return getPwdDTO, nil
}
