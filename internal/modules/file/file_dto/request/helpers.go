package request

import (
	"context"
	"fmt"
	"gophKeeper/internal/modules/auth/auth_middleware"
)

// GetUserIDFromContext Функция для извлечения userID из контекста
func getUserIDFromContext(ctx context.Context) (int, error) {
	if userID, ok := ctx.Value(auth_middleware.UserIDKey).(int); ok {
		return userID, nil // Возвращаем userID
	}
	return 0, fmt.Errorf("userID not found in context") // Ошибка, если не найден
}
