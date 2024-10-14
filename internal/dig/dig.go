package dig

import (
	"go.uber.org/dig"
)

var container *dig.Container

func InitDig() *dig.Container {
	if container == nil {
		return container
	}
	container = dig.New()
	return container
}
