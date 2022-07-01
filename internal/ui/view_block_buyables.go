package ui

import (
	"github.com/charmbracelet/lipgloss"
)

func DisplayBuyablesBlock(m Model) string {
	return block_buyables_style.
		Width(CalculateColumnWidth(m.Width, block_buyables_ratio)).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Left,
				DisplayUpgradesBlock(m),
				DisplayPrestigeBlock(m),
			))
}

func DisplayUpgradesBlock(m Model) string {
	upgrade_points_booster := "Buy upgrade: Increase points boost"

	return lipgloss.JoinVertical(
		lipgloss.Left,
		block_buyables_title_style.Render(upgrade_title),
		block_buyables_content_style.Render(upgrade_points_booster),
	)
}

func DisplayPrestigeBlock(m Model) string {
	column_size := CalculateColumnWidth(m.Width, block_buyables_ratio)
	prestige_points_booster := "Prestige to increase points boost"

	return lipgloss.JoinVertical(
		lipgloss.Left,
		block_buyables_title_style.Width(column_size).Render(prestige_title),
		block_buyables_content_style.Width(column_size).Render(prestige_points_booster),
	)
}
