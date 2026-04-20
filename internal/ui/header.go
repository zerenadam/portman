package ui

import "github.com/rivo/tview"

func (a *App) initHeader() {
	a.header = tview.NewTextView(). 
	SetText("portman - github.com/zerenadam/portman")
}
