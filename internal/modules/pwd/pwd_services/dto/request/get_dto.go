package request

import (
	"encoding/json"
	"fmt"
	"gophKeeper/internal/helpers"
	"io"
	"net/http"
)

type GetPwdDTO struct {
	PwdID  string `json:"pwd_id"`  // ID пароля в базе данных
	UserID int    `json:"user_id"` // Идентификатор пользователя
}

func GetPwdDTOFromHTTP(r *http.Request) (GetPwdDTO, error) {
	// Читаем тело запроса
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return GetPwdDTO{}, fmt.Errorf("reading body registration: %w", err)
	}
	defer r.Body.Close() // Закрываем тело запроса после чтения

	var getPwdDTO GetPwdDTO
	// Декодируем JSON в структуру
	err = json.Unmarshal(data, &getPwdDTO)
	if err != nil {
		return GetPwdDTO{}, fmt.Errorf("unmarshalling body registration: %w", err)
	}

	// Извлекаем userID из контекста
	userID, err := helpers.GetUserIDFromContext(r.Context())
	if err != nil {
		return GetPwdDTO{}, fmt.Errorf("error GetUserIDFromContext: %w", err)
	}

	getPwdDTO.UserID = userID // Устанавливаем userID в структуру
	return getPwdDTO, nil     // Возвращаем структуру и nil (без ошибки)
}
