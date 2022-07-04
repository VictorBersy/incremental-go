package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func DisplayModifiersBlock(m Model) string {
	return block_modifiers_style.
		Width(CalculateColumnWidth(m.Width, block_modifiers_ratio)).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Left,
				DisplayGeneratorsBlock(m),
				DisplayBoostersBlock(m),
			))
}

func DisplayGeneratorsBlock(m Model) string {
	column_size := CalculateColumnWidth(m.Width, block_modifiers_ratio)
	cost_generator := fmt.Sprintf("Cost: %0.1f", m.Resources.Points.GeneratorCost())
	generators_points := fmt.Sprintf("%s Points generator: %d (%.2f/s)", cost_generator, m.Resources.Points.Generators, m.Resources.Points.AddPerTick())

	return lipgloss.JoinVertical(
		lipgloss.Left,
		block_modifiers_title_style.Width(column_size).Render(generators_title),
		block_modifiers_content_style.Width(column_size).Render(generators_points),
	)
}

func DisplayBoostersBlock(m Model) string {
	column_size := CalculateColumnWidth(m.Width, block_modifiers_ratio)
	points_generation_boost := fmt.Sprintf("Points generation: x%0.1f", m.Resources.Points.Boost)
	containers_generations_boost := fmt.Sprintf("Containers generation: x%0.1f", m.Resources.Containers.Boost)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		block_modifiers_title_style.Width(column_size).Render(boosters_title),
		block_modifiers_content_style.Width(column_size).Render(points_generation_boost),
		block_modifiers_content_style.Width(column_size).Render(containers_generations_boost),
	)
}
