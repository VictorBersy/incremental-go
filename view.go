package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	s := fmt.Sprintln(lipgloss.NewStyle().Bold(true).Render("Whalecome! 🐳"))
	s += fmt.Sprintln(lipgloss.NewStyle().Bold(true).Render("Press 'P' to generate whale points"))
	if m.pointsGenerated > 25 {
		s += fmt.Sprintln(lipgloss.NewStyle().Bold(true).Render(fmt.Sprintf("Press 'A' to buy a container for %d points", m.config.costs["generators"]["points"])))
	}
	s += fmt.Sprintln(lipgloss.NewStyle().Render(fmt.Sprintf("Points created (total): %0.1f", m.pointsGenerated)))
	s += fmt.Sprintln(lipgloss.NewStyle().Render(fmt.Sprintf("Points created: %0.1f (generator: %d)", m.points, m.generator_points)))
	s += fmt.Sprintln(lipgloss.NewStyle().Render(fmt.Sprintf("Containers created: %0.1f (generator: %d)", m.containers, m.generator_containers)))
	if m.points > 0 && m.containers > 10 || m.pods > 0 {
		s += fmt.Sprintln(lipgloss.NewStyle().Render(fmt.Sprintf("Pods created: %0.1f", m.pods)))
	}
	if m.quitting {
		s += "\n"
	}
	return s
}
