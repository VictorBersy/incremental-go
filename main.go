package main

// A simple example that shows how to send activity to Bubble Tea in real-time
// through a channel.

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	sub                chan struct{} // where we'll receive activity notifications
	containers         uint64
	autobuy_containers uint64
	pods               uint64
	quitting           bool
}

// A message used to indicate that activity has occurred. In the real world (for
// example, chat) this would contain actual data.
type containerMsg struct{}

// Simulate a process that sends events at an irregular interval in real time.
// In this case, we'll send events on the channel at a random interval between
// 100 to 1000 milliseconds. As a command, Bubble Tea will run this
// asynchronously.
func listenForActivity(sub chan struct{}) tea.Cmd {
	return func() tea.Msg {
		for {
			time.Sleep(time.Millisecond * time.Duration(rand.Int63n(90)+10))
			sub <- struct{}{}
		}
	}
}

// A command that waits for the activity on a channel.
func waitForActivity(sub chan struct{}) tea.Cmd {
	return func() tea.Msg {
		return containerMsg(<-sub)
	}
}

func autobuy(m *model) {
	if m.autobuy_containers > 0 {
		ticker := time.NewTicker(10 * time.Millisecond)
		for range ticker.C {
			createContainer(m)
		}
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		listenForActivity(m.sub), // generate activity
		waitForActivity(m.sub),   // wait for activity
	)
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

	case containerMsg:
		return m, waitForActivity(m.sub) // wait for next event

	default:
		return m, nil
	}
}

func (m model) View() string {
	s := fmt.Sprintln(lipgloss.NewStyle().Bold(true).Render("Whalecome! 🐳"))
	if m.containers == 0 && m.pods == 0 {
		s += fmt.Sprintln(lipgloss.NewStyle().Bold(true).Render("Press 'A' to generate containers"))
	}
	s += fmt.Sprintln(lipgloss.NewStyle().Render(fmt.Sprintf("Containers created: %d (autobuyers: %d)", m.containers, m.autobuy_containers)))
	if m.containers > 10 || m.pods > 0 {
		s += fmt.Sprintln(lipgloss.NewStyle().Render(fmt.Sprintf("Pods created: %d", m.pods)))
	}
	if m.quitting {
		s += "\n"
	}
	return s
}

func main() {
	p := tea.NewProgram(model{
		sub: make(chan struct{}),
	},
		tea.WithAltScreen())

	if p.Start() != nil {
		fmt.Println("could not start program")
		os.Exit(1)
	}
}

func createContainer(m *model) model {
	m.containers++
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
