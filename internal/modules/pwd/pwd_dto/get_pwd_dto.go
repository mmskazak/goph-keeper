package pwd_dto

import (
	"context"
	"encoding/json"
	"fmt"
	"gophKeeper/internal/modules/auth/auth_middleware"
	"io"
	"net/http"
)

type GetPwdDTO struct {
	ID     string `json:"id"`      //ID пароля в базе данных
	UserID string `json:"user_id"` //Идентификатор пользователя
}

func GetPwdDTOFromHTTP(r *http.Request) (GetPwdDTO, error) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return GetPwdDTO{}, fmt.Errorf("reading body registration: %w", err)
	}
	var getPwdDTO GetPwdDTO
	err = json.Unmarshal(data, &getPwdDTO)
	if err != nil {
		return GetPwdDTO{}, fmt.Errorf("unmarshalling body registration: %w", err)
	}
	return getPwdDTO, nil
}

// GetUserIDFromContext Функция для извлечения userID из контекста
func GetUserIDFromContext(ctx context.Context) (int, error) {
	if userID, ok := ctx.Value(auth_middleware.UserIDKey).(int); ok {
		return userID, nil
	}
	return 0, fmt.Errorf("userID not found in context")
}
