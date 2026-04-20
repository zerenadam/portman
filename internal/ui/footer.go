package ui

import "github.com/rivo/tview"

func (a *App) initFooter() {
	a.footer = tview.NewTextView(). 
	SetText("[d] kill [r] reload [q] quit"). 
	SetTextAlign(tview.AlignCenter)
}
