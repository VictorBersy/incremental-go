package main

import "time"

type Config struct {
	tick       time.Duration
	costs      map[string]map[string]int
	generators map[string]float64
	boosters   map[string]float64
}

func LoadConfig() (config Config) {

	c := Config{
		tick: 10 * time.Millisecond,
		costs: map[string]map[string]int{
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
		generators: map[string]float64{
			"points_per_tick": 0.001,
		},
		boosters: map[string]float64{
			"points_per_tick": 0.001,
		},
	}

	return c
}
