package main

import "github.com/charmbracelet/lipgloss"

func generateLayout(m model) string {
	s := ""
	s += DisplayTopBar(m)
	s += lipgloss.JoinHorizontal(
		lipgloss.Top,
		DisplayTotalBlock(m),
		DisplayModifiersBlock(m),
		DisplayBuyersBlock(m),
	)

	return s
}
