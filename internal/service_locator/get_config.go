package service_locator

import (
	"fmt"
	"gophKeeper/internal/config"
)

func GetConfig() (*config.Config, error) {
	c := sc.Get("config")
	cfg, ok := c.(*config.Config)
	if !ok {
		return nil, fmt.Errorf("config is not config")
	}
	return cfg, nil
}
