package pwd_dto

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type SavePwdDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func GetSavePwdDTOFromHTTP(r *http.Request) (*SavePwdDTO, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("reading body registration: %w", err)
	}
	var savePwdDTO SavePwdDTO
	err = json.Unmarshal(data, &savePwdDTO)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling body registration: %w", err)
	}
	return &savePwdDTO, nil
}
