package components

import "github.com/rivo/tview"

func initFooter() *tview.TextView {
	footer := tview.NewTextView(). 
	SetText("[d] kill | [r] reload | [q] quit"). 
	SetTextAlign(tview.AlignCenter) 
	footer.SetBorder(true)
	return footer
}
