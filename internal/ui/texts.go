package ui

import (
	"fmt"

	"github.com/VictorBersy/incremental-go/internal/config"
)

// Top bar
var welcome_text string = fmt.Sprintf("Whalecome to incremental-go %s 🐳", config.GameVersion)
var goal_text string = "Reach XXX whale points"
var instructions_text string = "Press 'P' to generate your first points."

// Total block
var generated_title string = "Total generated"
var current_title string = "Current"

// Modifiers block
var generators_title string = "Generators"
var boosters_title string = "Boosters"

// Buyables block
var upgrade_title string = "Upgrades"
var prestige_title string = "Prestige"
