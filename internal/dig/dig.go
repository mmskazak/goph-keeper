package dig

import (
	"go.uber.org/dig"
)

var DI *dig.Container

func InitDig() *dig.Container {
	if DI == nil {
		return DI
	}
	DI = dig.New()
	return DI
}
