package ui

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/zerenadam/portman/internal/harbor"
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
		case 'd':
			a.killSelected()
		}
		return event 
	})
	
	a.populateTable()
}

func (a *App) populateTable() {
	a.table.Clear()
	
	headers := []string{"process", "PID", "address", "protocol", "user"}
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

func (a *App) killSelected() {
    row, _ := a.table.GetSelection()
    if row == 0 { return }  // don't kill on header row

    pidStr := a.table.GetCell(row, 1).Text
    pid, err := strconv.Atoi(pidStr)
    if err != nil { return }

    proc, err := os.FindProcess(pid)
    if err != nil {
        // show error somewhere
        return
    }

    if err := proc.Signal(os.Interrupt); err != nil {
        fmt.Println("kill failed:", err)
    }

    // refresh after kill
    a.ports = harbor.FindPorts()
    a.populateTable()
}

