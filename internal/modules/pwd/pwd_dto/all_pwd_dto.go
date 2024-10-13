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

func AllPwdDTOFromHTTP(r *http.Request) (DeletePwdDTO, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return DeletePwdDTO{}, fmt.Errorf("reading body registration: %w", err)
	}
	var deletePwdDTO DeletePwdDTO
	err = json.Unmarshal(data, &deletePwdDTO)
	if err != nil {
		return DeletePwdDTO{}, fmt.Errorf("unmarshalling body registration: %w", err)
	}
	return deletePwdDTO, nil
}
