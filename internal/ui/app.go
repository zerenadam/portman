package ui

import (
	"github.com/rivo/tview"
	"github.com/zerenadam/portman/internal/harbor"
)

type App struct {
	app *tview.Application
	layout *tview.Flex
	ports []*harbor.Port
}

func NewApp() *App {
	a := &App{
		app: tview.NewApplication(),
		layout: tview.NewFlex().SetDirection(tview.FlexRow),
		ports: harbor.FindPorts(),
	}
	return a
}

func (a *App) Run() error {
	return a.app.SetRoot(a.layout, true).Run()
}



