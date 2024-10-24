package authdto

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LoginDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func GetLoginDTOFromHTTP(r *http.Request) (*LoginDTO, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("reading body registration: %w", err)
	}
	var logDTO LoginDTO
	err = json.Unmarshal(data, &logDTO)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling body registration: %w", err)
	}
	return &logDTO, nil
}
