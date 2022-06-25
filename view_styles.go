package main

import (
	"github.com/charmbracelet/lipgloss"
)

func baseStyle(fgColor string, bgColor string) lipgloss.Style {
	return lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(fgColor)).
		Background(lipgloss.Color(bgColor)).
		Margin(1, 0, 0, 1).
		PaddingTop(2).
		PaddingBottom(2).
		PaddingLeft(4)
}

var welcome_style = baseStyle("ffffff", "#0a92be")

var total_style = baseStyle("ffffff", "#384d54")
var current_style = baseStyle("ffffff", "#384d54")

var modifiers_style = baseStyle("ffffff", "#e63946")
var buying_style = baseStyle("ffffff", "#2a9d8f")
