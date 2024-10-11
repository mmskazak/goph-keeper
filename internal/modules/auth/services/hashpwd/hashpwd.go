package hashpwd

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashAndStorePassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("could not generate hash: %v", err)
	}
	return string(bytes), nil
}

func CheckHashedPassword(hashedPassword, providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(providedPassword))
	return err == nil
}
