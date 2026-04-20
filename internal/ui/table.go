package ui

import "github.com/rivo/tview"

func (a *App) initTable() *tview.Table {
	table := tview.NewTable(). 
	SetSelectable(true, false). 
	SetFixed(1, 1)
	table.SetBorder(true)
	table.SetTitle(" ports ")
	
	headers := []string{"PROCESS", "PID", "ADDRESS", "PROTOCOL", "USER"} 

	for i, title := range headers {
		table.SetCell(0, i, tview.NewTableCell(title).SetExpansion(1))	
	}

	for row, port := range a.ports {
		r := row + 1
		table.SetCell(r, 0, tview.NewTableCell(port.Process).SetExpansion(1))
		table.SetCell(r, 1, tview.NewTableCell(port.PID).SetExpansion(1))
		table.SetCell(r, 2, tview.NewTableCell(port.Address).SetExpansion(1))
		table.SetCell(r, 3, tview.NewTableCell(port.Protocol).SetExpansion(1))
		table.SetCell(r, 4, tview.NewTableCell(port.User).SetExpansion(1))
	}
	return table
}
