package resources

import (
	"github.com/VictorBersy/incremental-go/internal/config"
)

type Resources struct {
	Containers Containers
	Pods       Pods
	Points     Points
}

func CalculatePerTick(perTick float64) float64 {
	return perTick * config.TickDuration().Seconds()
}
