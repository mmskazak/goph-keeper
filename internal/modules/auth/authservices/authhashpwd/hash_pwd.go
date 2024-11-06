package authhashpwd

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// GenerateHashFromPassword генерирует хеш из пароля.
func GenerateHashFromPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("could not generate hash: %w", err)
	}
	return string(bytes), nil
}

func IsEqualHashedPassword(hashedPassword, providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(providedPassword))
	return err == nil
}
