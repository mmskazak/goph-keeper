package dig

import (
	"fmt"
	"gophKeeper/internal/modules/pwd/pwd_http"
)

func RegisterPwdHandlers(pwd pwd_http.PwdHandlers) error {
	if err := container.Provide(pwd); err != nil {
		return fmt.Errorf("error creating dependency graph: %w", err)
	}
	return nil
}

func GetPwdHandlers() (*pwd_http.PwdHandlers, error) {
	var pwd pwd_http.PwdHandlers
	if err := container.Invoke(pwd); err != nil {
		return nil, fmt.Errorf("error invoking dependencies: %w", err)
	}
	return &pwd, nil
}
