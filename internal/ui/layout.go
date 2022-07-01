package ui

import (
	"github.com/charmbracelet/lipgloss"
)

var block_top_bar_ratio = float64(100)
var block_total_ratio = float64(20)
var block_modifiers_ratio = float64(40)
var block_buyables_ratio = float64(40)

func GenerateLayout(m Model) string {
	return lipgloss.JoinVertical(
		lipgloss.Top,
		DisplayTopBar(m),
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			DisplayTotalBlock(m),
			DisplayModifiersBlock(m),
			DisplayBuyablesBlock(m),
		),
	)
}
