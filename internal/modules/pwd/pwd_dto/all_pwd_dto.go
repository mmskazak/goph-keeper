package pwd_dto

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type AllPwdDTO struct {
	UserID string `json:"user_id"`
}

func AllPwdDTOFromHTTP(r *http.Request) (AllPwdDTO, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return AllPwdDTO{}, fmt.Errorf("reading body registration: %w", err)
	}
	var allPwdDTO AllPwdDTO
	err = json.Unmarshal(data, &allPwdDTO)
	if err != nil {
		return AllPwdDTO{}, fmt.Errorf("unmarshalling body registration: %w", err)
	}
	return allPwdDTO, nil
}
