package main

import (
	"fmt"
)

func DisplayModifiersBlock(m model) string {
	s := ""
	s += modifiers_style.Width(calculateColumnWidth(m.width, 10)).Underline(true).Render(generators_title)
	s += modifiers_style.Width(calculateColumnWidth(m.width, 10)).Underline(true).Render(boosters_title)
	return fmt.Sprintln(s)
}
