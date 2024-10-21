package helpers

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"gophKeeper/internal/modules/auth/auth_middleware"
)

// GetUserIDFromContext Функция для извлечения userID из контекста
func GetUserIDFromContext(ctx context.Context) (int, error) {
	claims, ok := ctx.Value(auth_middleware.Claims).(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("could not get claims from context")
	}

	fmt.Println(claims)

	// Приведение значения userID к float64, а затем к int
	userIDFloat, ok := claims["userID"].(float64)
	if !ok {
		return -1, fmt.Errorf("userID not found in claims or has incorrect type")
	}

	// Приведение float64 к int
	userID := int(userIDFloat)

	return userID, nil
}
