package request

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type AllPwdDTO struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Credentials string `json:"credentials"`
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
