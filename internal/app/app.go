package app

type App struct {
}

func (a *App) Init() {}

func (a *App) StartHTTP() {}
func (a *App) StopHTTP()  {}
func (a *App) StartGRPC() {}
func (a *App) StopGRPC()  {}
