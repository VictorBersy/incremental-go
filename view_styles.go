package main

import (
	"github.com/charmbracelet/lipgloss"
)

func blockStyle(fgColor string, bgColor string) lipgloss.Style {
	return lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(fgColor)).
		Background(lipgloss.Color(bgColor)).
		Margin(1, 0, 0, 1).
		PaddingTop(2).
		PaddingBottom(2).
		PaddingLeft(4)
}

var block_buyers_style = blockStyle("ffffff", "#2a9d8f")
var block_modifiers_style = blockStyle("ffffff", "#e63946")
var block_top_bar_style = blockStyle("ffffff", "#0a92be")
var block_total_style = blockStyle("ffffff", "#384d54")
