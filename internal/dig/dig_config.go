package dig

import (
	"fmt"
	"gophKeeper/internal/config"
)

func RegisterConfig(cfg config.Config) error {
	if err := container.Provide(cfg); err != nil {
		return fmt.Errorf("error creating dependency graph: %w", err)
	}
	return nil
}

func GetConfig() (*config.Config, error) {
	var cfg config.Config
	if err := container.Invoke(cfg); err != nil {
		return nil, fmt.Errorf("error invoking dependencies: %w", err)
	}
	return &cfg, nil

}
