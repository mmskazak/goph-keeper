package authdto

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type RegistrationDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func GetRegistrationDTOFromHTTP(r *http.Request) (*RegistrationDTO, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("reading body registration: %w", err)
	}
	var regDTO RegistrationDTO
	err = json.Unmarshal(data, &regDTO)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling body registration: %w", err)
	}
	return &regDTO, nil
}
