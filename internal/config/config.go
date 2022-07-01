package config

import "time"

type Config struct {
	Tick       time.Duration
	Costs      map[string]map[string]int
	Generators map[string]float64
	Boosters   map[string]float64
}

func Load() (config Config) {
	return Config{
		Tick: 10 * time.Millisecond,
		Costs: map[string]map[string]int{
			"models": {
				"container": 10,
				"pod":       25,
			},
			"generators": {
				"points":    10,
				"container": 25,
				"pod":       100,
			},
			"boosters": {
				"points":    1000,
				"container": 10000,
				"pod":       100000,
			},
		},
		Generators: map[string]float64{
			"points_per_tick": 0.001,
		},
		Boosters: map[string]float64{
			"points_per_tick": 0.001,
		},
	}
}
