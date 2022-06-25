package main

import (
	"fmt"
)

func DisplayTotalBlock(m model) string {
	s := ""
	s += total_style.Width(calculateColumnWidth(m.width, 20)).Underline(true).Render(current_title)
	s += total_style.Width(calculateColumnWidth(m.width, 20)).Underline(true).Render(generated_title)
	return fmt.Sprintln(s)
}
