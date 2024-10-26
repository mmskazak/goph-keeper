package request

import (
	"encoding/json"
	"fmt"
	"gophKeeper/internal/helpers"
	"gophKeeper/internal/logger"
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
		return GetPwdDTO{}, fmt.Errorf("reading body for GetPwdDTOFromHTTP: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Log.Errorf("error closing body reader: %v", err)
		}
	}(r.Body) // Закрываем тело запроса после чтения

	var getPwdDTO GetPwdDTO
	// Декодируем JSON в структуру
	err = json.Unmarshal(data, &getPwdDTO)
	if err != nil {
		return GetPwdDTO{}, fmt.Errorf("unmarshalling body registration: %w", err)
	}

	// Извлекаем userID из контекста
	userID, err := helpers.GetUserIDFromContext(r.Context())
	if err != nil {
		return GetPwdDTO{}, fmt.Errorf("error GetPwdDTOFromHTTP GetUserIDFromContext: %w", err)
	}

	getPwdDTO.UserID = userID // Устанавливаем userID в структуру
	return getPwdDTO, nil     // Возвращаем структуру и nil (без ошибки)
}
