package ui

import (
	"github.com/charmbracelet/lipgloss"
)

var block_buyables_colors = map[string]string{"fg": "ffffff", "bg": "#2a9d8f"}
var block_modifiers_colors = map[string]string{"fg": "ffffff", "bg": "#e63946"}
var block_top_bar_colors = map[string]string{"fg": "ffffff", "bg": "#0a92be"}
var block_total_colors = map[string]string{"fg": "ffffff", "bg": "#384d54"}

func blockStyle(fgColor string, bgColor string) lipgloss.Style {
	return lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(fgColor)).
		Background(lipgloss.Color(bgColor)).
		Height(25).
		Margin(1)
}

func contentStyle(fgColor string, bgColor string) lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color(fgColor)).
		Background(lipgloss.Color(bgColor)).
		Padding(0, 1, 0, 2)
}

func titleStyle(fgColor string, bgColor string) lipgloss.Style {
	return lipgloss.NewStyle().
		Bold(true).
		Underline(true).
		Foreground(lipgloss.Color(fgColor)).
		Background(lipgloss.Color(bgColor)).
		Padding(1)
}

// Blocks
var block_buyables_style = blockStyle(block_buyables_colors["fg"], block_buyables_colors["bg"])
var block_modifiers_style = blockStyle(block_modifiers_colors["fg"], block_modifiers_colors["bg"])
var block_top_bar_style = blockStyle(block_top_bar_colors["fg"], block_top_bar_colors["bg"]).Padding(2, 5)
var block_total_style = blockStyle(block_total_colors["fg"], block_total_colors["bg"])

// Title
var block_buyables_title_style = titleStyle(block_buyables_colors["fg"], block_buyables_colors["bg"])
var block_modifiers_title_style = titleStyle(block_modifiers_colors["fg"], block_modifiers_colors["bg"])
var block_total_title_style = titleStyle(block_total_colors["fg"], block_total_colors["bg"])

// Content
var block_buyables_content_style = contentStyle(block_buyables_colors["fg"], block_buyables_colors["bg"])
var block_modifiers_content_style = contentStyle(block_modifiers_colors["fg"], block_modifiers_colors["bg"])
var block_total_content_style = contentStyle(block_total_colors["fg"], block_total_colors["bg"])
