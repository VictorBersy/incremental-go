package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func DisplayBuyablesBlock(m model) string {
	return block_buyables_style.
		Width(calculateColumnWidth(m.width, block_buyables_ratio)).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Left,
				DisplayUpgradesBlock(m),
				DisplayPrestigeBlock(m),
			))
}

func DisplayUpgradesBlock(m model) string {
	upgrade_points_booster := fmt.Sprintf("Buy upgrade: Increase points boost")

	return lipgloss.JoinVertical(
		lipgloss.Left,
		block_buyables_title_style.Render(upgrade_title),
		block_buyables_content_style.Render(upgrade_points_booster),
	)
}

func DisplayPrestigeBlock(m model) string {
	column_size := calculateColumnWidth(m.width, block_buyables_ratio)
	prestige_points_booster := fmt.Sprintf("Prestige to increase points boost")

	return lipgloss.JoinVertical(
		lipgloss.Left,
		block_buyables_title_style.Width(column_size).Render(prestige_title),
		block_buyables_content_style.Width(column_size).Render(prestige_points_booster),
	)
}
