package ui

import "github.com/rivo/tview"

func (a *App) initFooter() *tview.TextView {
	footer := tview.NewTextView(). 
	SetText("[d] kill | [r] reload | [q] quit"). 
	SetTextAlign(tview.AlignCenter) 
	footer.SetBorder(true)
	return footer
}

