package request

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type DeletePwdDTO struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Credentials string `json:"credentials"`
}

func DeletePwdDTOFromHTTP(r *http.Request) (DeletePwdDTO, error) {
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
