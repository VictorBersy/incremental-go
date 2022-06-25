package main

import (
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	width  int
	height int

	quitting bool
	config   Config

	points          float64
	pointsGenerated float64
	pointsBoost     float64
	pointsGenerator int

	containers          float64
	containersGenerated float64
	containersBoost     float64
	containersGenerator int

	pods          float64
	podsGenerated float64
	podsBoost     float64
}

type tickMsg time.Time

func main() {
	p := tea.NewProgram(model{config: LoadConfig()}, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(tick(&m), tea.EnterAltScreen)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "p":
			addPoint(&m, 1.0)
			return m, nil
		case "o":
			m = buyPointGenerator(&m)
			return m, nil
		case "a":
			buyContainer(&m)
			return m, nil
		case "b":
			m = buyPods(&m)
			return m, nil
		case "c":
			m = buyContainerGenerator(&m)
			return m, nil
		case "esc", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		default:
			return m, nil
		}

	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		return m, nil

	case tickMsg:
		if m.pointsGenerator > 0 {
			generatePoint(&m)
		}
		return m, tick(&m)

	default:
		return m, nil
	}
}

func tick(m *model) tea.Cmd {
	return tea.Tick(m.config.tick, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
