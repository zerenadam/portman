package ui

import (
	"github.com/rivo/tview"
	"github.com/zerenadam/portman/internal/harbor"
)

type App struct {
	app *tview.Application
	layout *tview.Flex
	header *tview.TextView
	table *tview.Table
	footer *tview.TextView
	ports []*harbor.Port
}

func NewApp() *App {
	a := &App{
		app: tview.NewApplication(),
		layout: tview.NewFlex().SetDirection(tview.FlexRow),
		ports: harbor.FindPorts(),
	}
	a.build()
	return a 
}

func (a *App) Run() error {
	return a.app.SetRoot(a.layout, true).Run()
}

func (a *App) build() {
	a.initHeader()
	a.initTable()
	a.initFooter()

	a.layout.AddItem(a.header, 1, 1, false). 
	AddItem(a.table, 0, 1, true). 
	AddItem(a.footer, 3, 1, false)
}
