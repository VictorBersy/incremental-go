package ui

import (
	"time"

	"github.com/VictorBersy/incremental-go/internal/config"
	"github.com/VictorBersy/incremental-go/internal/resources"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Width  int
	Height int

	tea.Model

	Quitting  bool
	Resources resources.Resources
}

type TickMsg time.Time

func NewModel() Model {
	return Model{
		Width:    0,
		Height:   0,
		Model:    nil,
		Quitting: false,
		Config:   config.Load(),
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(tick(m.Config.Tick), tea.EnterAltScreen)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "p":
			m.Resources.Points = m.Resources.Points.ManuallyGenerate()
			return m, nil
		case "o":
			m.Resources.Points = m.Resources.Points.BuyGenerator()
			return m, nil
		case "a":
			m.Resources = m.Resources.Containers.Buy(m.Resources)
			return m, nil
		case "b":
			m.Resources.Pods = m.Resources.Pods.Buy()
			return m, nil
		case "c":
			m.Resources.Containers = m.Resources.Containers.BuyGenerator()
			return m, nil
		case "esc", "ctrl+c":
			m.Quitting = true
			return m, tea.Quit
		default:
			return m, nil
		}

	case tea.WindowSizeMsg:
		m.Width, m.Height = msg.Width, msg.Height
		return m, nil

	case TickMsg:
		m.Resources.Points = m.Resources.Points.Generate()
		return m, tick(m.Config.Tick)

	default:
		return m, nil
	}
}

func (m Model) View() string {
	s := GenerateLayout(m)
	if m.Quitting {
		s += "\n"
	}
	return s
}

func tick(duration time.Duration) tea.Cmd {
	return tea.Tick(duration, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}
