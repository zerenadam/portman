package ui

import "github.com/rivo/tview"

func (a *App) initHeader() *tview.TextView {
	logo := `
                      __                      
    ____  ____  _____/ /_____ ___  ____ _____ 
   / __ \/ __ \/ ___/ __/ __ ` + "`" + `__ \/ __ ` + "`" + `/ __ \
  / /_/ / /_/ / /  / /_/ / / / / / /_/ / / / /
 / .___/\____/_/   \__/_/ /_/ /_/\__,_/_/ /_/ 
/_/                                           
	`	
	header := tview.NewTextView(). 
	SetText(logo). 
	SetWrap(true). 
	SetTextAlign(tview.AlignCenter)
	header.SetBorder(true)

	return header
}


