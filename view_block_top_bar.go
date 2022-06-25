package main

import (
	"fmt"
)

func DisplayTopBar(m model) string {
	s := ""
	s += block_top_bar_style.Width(m.width).Render(welcome_text)
	s += block_top_bar_style.Width(m.width).Render(goal_text)
	s += block_top_bar_style.Width(m.width).Render(instructions_text)
	return fmt.Sprintln(s)
}
