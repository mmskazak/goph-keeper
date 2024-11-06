package helpers

import (
	"context"
	"errors"
	"fmt"
	"goph-keeper/internal/modules/auth/authmiddleware"
	"goph-keeper/internal/modules/auth/authservices/authjwtservice"

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

// ExtractUserIDFromToken извлекает userID из JWT токена и приводит его к int.
func ExtractUserIDFromToken(token *jwt.Token) (int, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.New("invalid JWT claims")
	}

	// Извлечение и проверка userID
	userIDFloat, ok := claims["userID"].(float64)
	if !ok {
		return 0, errors.New("error extracting userID from claims")
	}

	return int(userIDFloat), nil
}

// ParseTokenAndExtractUserID выполняет парсинг и валидацию JWT токена, а затем извлекает userID.
func ParseTokenAndExtractUserID(jwtToken, secretKey string) (int, error) {
	token, err := authjwtservice.ParseAndValidateToken(jwtToken, secretKey)
	if err != nil {
		return 0, fmt.Errorf("error parsing and validating JWT token: %w", err)
	}

	userID, err := ExtractUserIDFromToken(token)
	if err != nil {
		return 0, fmt.Errorf("error extracting userID: %w", err)
	}

	return userID, nil
}
