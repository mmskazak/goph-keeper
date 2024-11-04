package authdto

import (
	"encoding/json"
	"errors"
	"fmt"
	pb "goph-keeper/internal/modules/auth/proto"
	"io"
	"net/http"
)

type RegistrationDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetRegistrationDTOFromHTTP(r *http.Request) (*RegistrationDTO, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("reading body for get registration dto: %w", err)
	}
	var regDTO RegistrationDTO
	err = json.Unmarshal(data, &regDTO)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling body registration: %w", err)
	}
	return &regDTO, nil
}

// GetRegistrationDTOFromRegistrationRequestGRPC преобразует RegistrationRequest в RegistrationDTO
func GetRegistrationDTOFromRegistrationRequestGRPC(req *pb.RegistrationRequest) (*RegistrationDTO, error) {
	// Проверяем, что логин и пароль не пустые
	if req.GetUsername() == "" || req.GetPassword() == "" {
		return nil, errors.New("login and password must not be empty")
	}

	return &RegistrationDTO{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
	}, nil
}
