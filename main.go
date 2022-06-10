package main

import (
	"fmt"
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	containers         float64
	pods               float64
	autobuy_containers uint64
	quitting           bool
}

type tickMsg time.Time

func main() {
	p := tea.NewProgram(model{}, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(tick(), tea.EnterAltScreen)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "a":
			createContainer(&m)
			return m, nil
		case "b":
			m = buyPods(m)
			return m, nil
		case "c":
			m = buyContainerGenerator(m)
			return m, nil
		case "esc", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		default:
			return m, nil
		}

	case tickMsg:
		if m.autobuy_containers > 0 {
			autoBuyContainer(&m)
		}
		return m, tick()

	default:
		return m, nil
	}
}

func (m model) View() string {
	s := fmt.Sprintln(lipgloss.NewStyle().Bold(true).Render("Whalecome! 🐳"))
	if m.containers == 0 && m.pods == 0 {
		s += fmt.Sprintln(lipgloss.NewStyle().Bold(true).Render("Press 'A' to generate containers"))
	}
	s += fmt.Sprintln(lipgloss.NewStyle().Render(fmt.Sprintf("Containers created: %0.2f (autobuyers: %d)", m.containers, m.autobuy_containers)))
	if m.containers > 10 || m.pods > 0 {
		s += fmt.Sprintln(lipgloss.NewStyle().Render(fmt.Sprintf("Pods created: %0.2f", m.pods)))
	}
	if m.quitting {
		s += "\n"
	}
	return s
}

func createContainer(m *model) model {
	m.containers++
	return *m
}

func autoBuyContainer(m *model) model {
	m.containers = m.containers + (float64(1*m.autobuy_containers) / 10)
	return *m
}

func buyContainerGenerator(m model) model {
	if m.containers >= 25 {
		m.containers = m.containers - 25
		m.autobuy_containers++
	}
	return m
}

func buyPods(m model) model {
	if m.containers >= 100 {
		m.containers = m.containers - 100
		m.pods++
	}
	return m
}

func tick() tea.Cmd {
	return tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
