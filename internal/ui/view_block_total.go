package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func DisplayTotalBlock(m Model) string {
	return block_total_style.
		Width(CalculateColumnWidth(m.Width, block_total_ratio)).
		Render(
			lipgloss.JoinVertical(
				lipgloss.Left,
				DisplayGeneratedBlock(m),
				DisplayCurrentBlock(m),
			))
}

func DisplayGeneratedBlock(m Model) string {
	column_size := CalculateColumnWidth(m.Width, block_total_ratio)
	points := fmt.Sprintf("Points: %0.1f", m.Resources.Points.Generated)
	containers := fmt.Sprintf("Containers: %0.1f", m.Resources.Containers.Generated)
	pods := fmt.Sprintf("Pods: %0.1f", m.Resources.Pods.Generated)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		block_total_title_style.Width(column_size).Render(generated_title),
		block_total_content_style.Width(column_size).Render(points),
		block_total_content_style.Width(column_size).Render(containers),
		block_total_content_style.Width(column_size).Render(pods),
	)
}

func DisplayCurrentBlock(m Model) string {
	column_size := CalculateColumnWidth(m.Width, block_total_ratio)
	points := fmt.Sprintf("Points: %0.1f", m.Resources.Points.Count)
	containers := fmt.Sprintf("Containers: %0.1f", m.Resources.Containers.Count)
	pods := fmt.Sprintf("Pods: %0.1f", m.Resources.Pods.Count)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		block_total_title_style.Width(column_size).Render(current_title),
		block_total_content_style.Width(column_size).Render(points),
		block_total_content_style.Width(column_size).Render(containers),
		block_total_content_style.Width(column_size).Render(pods),
	)
}
