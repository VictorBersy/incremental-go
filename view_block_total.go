package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func DisplayTotalBlock(m model) string {
	return block_total_style.Width(calculateColumnWidth(m.width, block_total_ratio)).Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			DisplayGeneratedBlock(m),
			DisplayCurrentBlock(m),
		))
}

func DisplayGeneratedBlock(m model) string {
	column_size := calculateColumnWidth(m.width, block_total_ratio)
	points := fmt.Sprintf("Points: %0.1f", m.pointsGenerated)
	containers := fmt.Sprintf("Containers: %0.1f", m.containersGenerated)
	pods := fmt.Sprintf("Pods: %0.1f", m.podsGenerated)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		block_total_title_style.Width(column_size).Render(generated_title),
		block_total_content_style.Width(column_size).Render(points),
		block_total_content_style.Width(column_size).Render(containers),
		block_total_content_style.Width(column_size).Render(pods),
	)
}

func DisplayCurrentBlock(m model) string {
	column_size := calculateColumnWidth(m.width, block_total_ratio)
	points := fmt.Sprintf("Points: %0.1f", m.points)
	containers := fmt.Sprintf("Containers: %0.1f", m.containers)
	pods := fmt.Sprintf("Pods: %0.1f", m.pods)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		block_total_title_style.Width(column_size).Render(current_title),
		block_total_content_style.Width(column_size).Render(points),
		block_total_content_style.Width(column_size).Render(containers),
		block_total_content_style.Width(column_size).Render(pods),
	)
}
