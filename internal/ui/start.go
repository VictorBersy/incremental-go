package ui

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func Start() {
	p := tea.NewProgram(NewModel())

	p.EnableMouseCellMotion()
	p.EnterAltScreen()
	err := p.Start()
	if err != nil {
		log.Fatal(err)
		p.ExitAltScreen()
		p.DisableMouseCellMotion()
	}
}
