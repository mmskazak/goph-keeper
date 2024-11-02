package authdto

import (
	"encoding/json"
	"fmt"
	pb "goph-keeper/internal/modules/auth/proto"
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
	if req.Login == "" || req.Password == "" {
		return nil, fmt.Errorf("login and password must not be empty")
	}

	return &RegistrationDTO{
		Login:    req.Login,
		Password: req.Password,
	}, nil
}
