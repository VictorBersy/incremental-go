package main

import (
	"fmt"
)

func DisplayTopBar(m model) string {
	s := ""
	s += welcome_style.Width(m.width).Render(welcome_text)
	s += welcome_style.Width(m.width).Render(goal_text)
	s += welcome_style.Width(m.width).Render(instructions_text)
	return fmt.Sprintln(s)
}
