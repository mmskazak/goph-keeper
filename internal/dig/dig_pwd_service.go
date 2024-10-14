package dig

import (
	"fmt"
	"gophKeeper/internal/modules/pwd/pwd_services"
)

func RegisterPwdService(pwd pwd_services.PwdService) error {
	if err := container.Provide(pwd); err != nil {
		return fmt.Errorf("error creating dependency graph: %w", err)
	}
	return nil
}

func GetPwdService() (*pwd_services.PwdService, error) {
	var auth pwd_services.PwdService
	if err := container.Invoke(auth); err != nil {
		return nil, fmt.Errorf("error invoking dependencies: %w", err)
	}
	return &auth, nil
}
