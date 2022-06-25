package main

import (
	"fmt"
)

func DisplayTotalBlock(m model) string {
	s := ""
	s += block_total_style.Width(calculateColumnWidth(m.width, 20)).Underline(true).Render(current_title)
	s += block_total_style.Width(calculateColumnWidth(m.width, 20)).Underline(true).Render(generated_title)
	return fmt.Sprintln(s)
}
