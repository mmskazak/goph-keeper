package request

import (
	"encoding/json"
	"fmt"
	"gophKeeper/internal/modules/pwd/pwd_services/value_obj"
	"io"
	"net/http"
)

type SaveFileDTO struct {
	UserID      int                   `json:"user_id"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	Credentials value_obj.Credentials `json:"credentials"`
}

func SaveFileDTOFromHTTP(r *http.Request) (SaveFileDTO, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return SaveFileDTO{}, fmt.Errorf("reading body registration: %w", err)
	}
	var savePwdDTO SaveFileDTO
	err = json.Unmarshal(data, &savePwdDTO)
	if err != nil {
		return SaveFileDTO{}, fmt.Errorf("unmarshalling body registration: %w", err)
	}

	// Извлекаем userID из контекста
	userID, err := getUserIDFromContext(r.Context())
	if err != nil {
		return SaveFileDTO{}, fmt.Errorf("error getUserIDFromContext: %w", err)
	}

	savePwdDTO.UserID = userID

	return savePwdDTO, nil
}
