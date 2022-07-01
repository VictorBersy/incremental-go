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
	Config    config.Config
	Resources struct {
		Containers resources.Containers
		Pods       resources.Pods
		Points     resources.Points
	}
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
	return tea.Batch(tick(&m), tea.EnterAltScreen)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "p":
			m.Resources.Points = m.Resources.Points.Add(1.0)
			return m, nil
		case "o":
			m.Resources.Points = m.Resources.Points.BuyGenerator()
			return m, nil
		case "a":
			m.Resources.Containers = m.Resources.Containers.Buy()
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
		if m.Resources.Points.Generators > 0 {
			m.Resources.Points.Generate()
		}
		return m, tick(&m)

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

func tick(m *Model) tea.Cmd {
	return tea.Tick(m.Config.Tick, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}
