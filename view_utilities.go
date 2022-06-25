package main

import (
	"math"
)

func calculateColumnWidth(width int, ratio float64) int {
	return int(math.Ceil(float64(width) * (ratio / 100)))
}
