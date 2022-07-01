package ui

import (
	"fmt"
)

func DisplayTopBar(m Model) string {
	text := fmt.Sprintln(welcome_text) +
		fmt.Sprintln(goal_text) +
		fmt.Sprintln(instructions_text)

	return block_top_bar_style.
		Width(CalculateColumnWidth(m.Width, block_top_bar_ratio)).
		Height(5).
		Render(text)
}
