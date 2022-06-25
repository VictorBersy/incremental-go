package main

import "fmt"

func DisplayTopBar(m model) string {
	text := fmt.Sprintln(welcome_text) + fmt.Sprintln(goal_text) + fmt.Sprintln(instructions_text)
	return block_top_bar_style.Width(calculateColumnWidth(m.width, block_top_bar_ratio)).Height(5).Render(text)
}
