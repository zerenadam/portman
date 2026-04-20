package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (a *App) initTable() {
	a.table = tview.NewTable().SetFixed(1, 0)
	a.table.SetTitle("ports")
	a.table.SetBorder(true)
	a.table.SetSelectable(true, false)
	a.table.SetFixed(1, 0)

	a.table.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'q':
			a.app.Stop()
			
		}
		return event 
	})
	
	a.populateTable()
}

func (a *App) populateTable() {
	a.table.Clear()
	
	headers := []string{"process", "PID", "@", "protocol", "user"}
	for col, h := range headers {
		a.table.SetCell(0, col, tview.NewTableCell(h).
		SetExpansion(1).
		SetSelectable(false))
	}

	for row, p := range a.ports {
		r := row + 1
		a.table.SetCell(r, 0, tview.NewTableCell(p.Process).SetExpansion(1))
		a.table.SetCell(r, 1, tview.NewTableCell(p.PID).SetExpansion(1))
		a.table.SetCell(r, 2, tview.NewTableCell(p.Address).SetExpansion(1))
		a.table.SetCell(r, 3, tview.NewTableCell(p.Protocol).SetExpansion(1))
		a.table.SetCell(r, 4, tview.NewTableCell(p.User).SetExpansion(1))
	}
}
