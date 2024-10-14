package dig

import (
	"fmt"
	"gophKeeper/internal/modules/auth/auth_services/auth_service"
)

func RegisterAuthService(auth auth_service.AuthService) error {
	if err := container.Provide(auth); err != nil {
		return fmt.Errorf("error creating dependency graph: %w", err)
	}
	return nil
}

func GetAuthService() (*auth_service.AuthService, error) {
	var auth auth_service.AuthService
	if err := container.Invoke(auth); err != nil {
		return nil, fmt.Errorf("error invoking dependencies: %w", err)
	}
	return &auth, nil
}
