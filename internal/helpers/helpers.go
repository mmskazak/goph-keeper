package helpers

import (
	"context"
	"errors"
	"fmt"
	"gophKeeper/internal/modules/auth/authmiddleware"

	"github.com/golang-jwt/jwt"
)

// GetUserIDFromContext Функция для извлечения userID из контекста.
func GetUserIDFromContext(ctx context.Context) (int, error) {
	claims, ok := ctx.Value(authmiddleware.Claims).(jwt.MapClaims)
	if !ok {
		return 0, errors.New("no claims found in context")
	}

	fmt.Println(claims)

	// Приведение значения userID к float64, а затем к int
	userIDFloat, ok := claims["userID"].(float64)
	if !ok {
		return -1, errors.New("no userID found in context")
	}

	// Приведение float64 к int
	userID := int(userIDFloat)

	return userID, nil
}
