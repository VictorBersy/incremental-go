package main

import (
	"fmt"
)

func DisplayModifiersBlock(m model) string {
	s := ""
	s += block_modifiers_style.Width(calculateColumnWidth(m.width, 10)).Underline(true).Render(generators_title)
	s += block_modifiers_style.Width(calculateColumnWidth(m.width, 10)).Underline(true).Render(boosters_title)
	return fmt.Sprintln(s)
}
