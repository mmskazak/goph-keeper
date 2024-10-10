package app

import (
	"fmt"
	"gophKeeper/internal/config"
	"gophKeeper/internal/modules/auth/http"
)

type ServiceLocator struct {
	container map[string]any
}

var sc *ServiceLocator

// InitServiceLocator Синглтон
func InitServiceLocator() *ServiceLocator {
	if sc == nil {
		sc = &ServiceLocator{
			container: make(map[string]any),
		}
	}
	return sc
}

func (sc *ServiceLocator) Register(name string, service any) {
	sc.container[name] = service
}

func (sc *ServiceLocator) Get(name string) any {
	return sc.container[name]
}

func GetConfig() (*config.Config, error) {
	c := sc.Get("config")
	cfg, ok := c.(*config.Config)
	if !ok {
		return nil, fmt.Errorf("config is not config")
	}
	return cfg, nil
}

func GetAuthHandlersHTTP() (auth_http.AuthHandlers, error) {
	handlers := sc.Get("auth_handlers_http")
	ah, ok := handlers.(auth_http.AuthHandlers)
	if !ok {
		return auth_http.AuthHandlers{}, fmt.Errorf("auth_handlers is not auth_http.AuthHandlers")
	}
	return ah, nil
}
