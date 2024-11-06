package authdto

import (
	"encoding/json"
	"errors"
	"fmt"
	pb "goph-keeper/internal/modules/auth/proto"

	"io"
	"net/http"
)

// LoginDTO ...
type LoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginDTOFromRequestHTTP преобразует http запрос в LoginDTO.
func LoginDTOFromRequestHTTP(r *http.Request) (*LoginDTO, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("reading body for get login dto: %w", err)
	}
	var logDTO LoginDTO
	err = json.Unmarshal(data, &logDTO)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling body registration: %w", err)
	}
	return &logDTO, nil
}

// LoginDTOFromLoginRequestGRPC преобразует LoginRequest в LoginDTO.
func LoginDTOFromLoginRequestGRPC(req *pb.LoginRequest) (*LoginDTO, error) {
	if req.GetUsername().GetUsername().GetValue() == "" || req.GetPassword().GetPassword().GetValue() == "" {
		return nil, errors.New("username and password must not be empty")
	}

	return &LoginDTO{
		Username: req.GetUsername().GetUsername().GetValue(),
		Password: req.GetPassword().GetPassword().GetValue(),
	}, nil
}
