package service_locator

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
