package dig

import (
	"fmt"
	"gophKeeper/internal/modules/auth/auth_http"
)

func RegisterAuthHandlers(auth auth_http.AuthHandlers) error {
	if err := container.Provide(auth); err != nil {
		return fmt.Errorf("error creating dependency graph: %w", err)
	}
	return nil
}

func GetAuthHandlers() (*auth_http.AuthHandlers, error) {
	var auth auth_http.AuthHandlers
	if err := container.Invoke(auth); err != nil {
		return nil, fmt.Errorf("error invoking dependencies: %w", err)
	}
	return &auth, nil
}
