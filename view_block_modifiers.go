package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func DisplayModifiersBlock(m model) string {
	return block_modifiers_style.
		Width(calculateColumnWidth(m.width, block_modifiers_ratio)).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Left,
				DisplayGeneratorsBlock(m),
				DisplayBoostersBlock(m),
			))
}

func DisplayGeneratorsBlock(m model) string {
	column_size := calculateColumnWidth(m.width, block_modifiers_ratio)
	generators_points := fmt.Sprintf("Points generator: %d/s", m.pointsGenerator)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		block_modifiers_title_style.Width(column_size).Render(generators_title),
		block_modifiers_content_style.Width(column_size).Render(generators_points),
	)
}

func DisplayBoostersBlock(m model) string {
	column_size := calculateColumnWidth(m.width, block_modifiers_ratio)
	points_generation_boost := fmt.Sprintf("Points generation: x%0.1f", m.pointsBoost)
	containers_generations_boost := fmt.Sprintf("Containers generation: x%0.1f", m.containersBoost)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		block_modifiers_title_style.Width(column_size).Render(boosters_title),
		block_modifiers_content_style.Width(column_size).Render(points_generation_boost),
		block_modifiers_content_style.Width(column_size).Render(containers_generations_boost),
	)
}
